package templates

import (
	"fmt"

	"jamf/helper"
)

// Model - Template for models.
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

// Interface - Template to generate model's interface.
func Interface(name, face string) string {
	tmpl := `import { Document } from "mongoose";
	
export interface %s extends Document {%s}
`

	return fmt.Sprintf(tmpl, helper.Capitalize(name), face)
}
