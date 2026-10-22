package admin

func DeleteCalendarResource(conf map[string]string, key string) error {
	zcs, err := Authenticate(conf)

	if err != nil {
		return err
	}

	id, err := getCalendarResourceId(zcs, key)

	if err != nil {
		return err
	}

	return zcs.DeleteCalendarResourceRequest(id)
}
