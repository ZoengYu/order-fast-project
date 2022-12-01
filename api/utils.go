package api

func calculate_offset(page_id int32, page_size int32) int32 {
	return (page_id - 1) * page_size
}
