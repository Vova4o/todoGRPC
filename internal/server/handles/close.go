package handles

func (h *Handlers) Close() string {
	text := h.serviceLevel.Close()
	return text
}
