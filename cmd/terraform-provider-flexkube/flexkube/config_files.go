package flexkube

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func configFilesMarshal(c map[string]string, sensitive bool) interface{} {
	i := map[string]interface{}{}

	for k, v := range c {
		if v == "" {
			continue
		}

		if !sensitive {
			i[k] = v

			continue
		}

		i[k] = sha256sum([]byte(v))
	}

	return i
}

func configFilesUnmarshal(i interface{}) map[string]string {
	cf := map[string]string{}

	if i == nil {
		return cf
	}

	for k, v := range i.(map[string]interface{}) {
		cf[k] = v.(string)
	}

	return cf
}

func configFilesSchema(computed bool) *schema.Schema {
	return optionalMapPrimitive(computed, func(computed bool) *schema.Schema {
		return &schema.Schema{
			Type: schema.TypeString,
		}
	})
}