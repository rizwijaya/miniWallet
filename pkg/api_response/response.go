package api

func CustomResponse(data interface{}, status string) Response {
	response := Response{
		Status: status,
	}

	if status == HttpStatusSuccess {
		response.Data = data
	} else if status == HttpStatusFailed {
		response.Data = map[string]interface{}{
			"error": data,
		}
	}

	return response
}
