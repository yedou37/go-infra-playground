package middleware

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func okHandler(body string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
	})
}

func panickingHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("kaboom")
	})
}

func doRequest(t *testing.T, h http.Handler, method, path, token string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(method, path, nil)
	if token != "" {
		req.Header.Set("X-Token", token)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestChainOrder(t *testing.T) {
	var sink []string
	first := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sink = append(sink, "first-pre")
			next.ServeHTTP(w, r)
			sink = append(sink, "first-post")
		})
	}
	second := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sink = append(sink, "second-pre")
			next.ServeHTTP(w, r)
			sink = append(sink, "second-post")
		})
	}
	core := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sink = append(sink, "core")
	})

	h := Chain(first, second)(core)
	doRequest(t, h, http.MethodGet, "/", "")
	want := []string{"first-pre", "second-pre", "core", "second-post", "first-post"}
	if !reflect.DeepEqual(sink, want) {
		t.Fatalf("middleware order wrong\n got %v\nwant %v", sink, want)
	}
}

func TestLogMiddlewareRecordsRequest(t *testing.T) {
	var sink []string
	h := LogMiddleware(&sink)(okHandler("ok"))
	doRequest(t, h, http.MethodGet, "/foo", "")
	doRequest(t, h, http.MethodPost, "/bar", "")
	want := []string{"GET /foo", "POST /bar"}
	if !reflect.DeepEqual(sink, want) {
		t.Fatalf("log got %v, want %v", sink, want)
	}
}

func TestAuthMiddlewareRejectsBadToken(t *testing.T) {
	h := AuthMiddleware("good")(okHandler("ok"))
	rec := doRequest(t, h, http.MethodGet, "/", "")
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("missing token: code=%d, want 401", rec.Code)
	}
	rec = doRequest(t, h, http.MethodGet, "/", "bad")
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("bad token: code=%d, want 401", rec.Code)
	}
	rec = doRequest(t, h, http.MethodGet, "/", "good")
	if rec.Code != http.StatusOK {
		t.Fatalf("good token: code=%d, want 200", rec.Code)
	}
}

func TestRecoverMiddlewareTurnsPanicInto500(t *testing.T) {
	h := RecoverMiddleware()(panickingHandler())
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("RecoverMiddleware must not let panic escape; recovered %v", r)
		}
	}()
	rec := doRequest(t, h, http.MethodGet, "/", "")
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("code=%d, want 500", rec.Code)
	}
}

func TestChainCombinesAuthLogRecover(t *testing.T) {
	var sink []string
	h := Chain(
		RecoverMiddleware(),
		LogMiddleware(&sink),
		AuthMiddleware("good"),
	)(panickingHandler())

	rec := doRequest(t, h, http.MethodGet, "/protected", "good")
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("code=%d, want 500 (recover should fire even when handler panics)", rec.Code)
	}
	if len(sink) != 1 || sink[0] != "GET /protected" {
		t.Fatalf("LogMiddleware should still record the authed request; got %v", sink)
	}
}
