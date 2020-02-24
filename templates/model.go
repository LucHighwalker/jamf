package templates

import (
	"fmt"

	"jamf/helper"
)

func Model(name, fields string) string {
	tmpl := `import { Schema, Model, model } from "mongoose";
import { %s } from "../interfaces/%s";

export const %sSchema: Schema = new Schema({%s});

const %sModel: Model<%s> = model<%s>("%s", %sSchema);

export default %sModel;
`

	cn := helper.Capitalize(name)

	return fmt.Sprintf(tmpl, cn, name, cn, fields, cn, cn, cn, cn, cn, cn)
}
