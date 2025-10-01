package http_controller

import "net/http"

func (c *Controller) Resolve(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	if code == "" {
		http.NotFound(w, r)
		return
	}

	orig, err := c.uc.Resolve(code)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, orig, http.StatusFound)
}
