package gomodifytags

import (
	"fmt"
	"os"
	"strings"
)

func newDefaultConfig() *config {
	return &config{
		file:                 "", //Filename to be parsed
		line:                 "", //Line number of the field or a range of line. i.e: 4 or 4,
		structName:           "", //Struct name to be processed
		fieldName:            "", //Field name to be processed
		fieldSeparator:       "|",
		offset:               0,           //Byte offset of the cursor position inside a struct.Can be anwhere from the comment until closing bracket
		all:                  false,       //Select all structs to be processed
		output:               "source",    //Output format. By default it's the whole file. Options: [source, json]
		write:                false,       //Write result to (source) file instead of stdout
		clear:                false,       //Clear all tags
		clearOption:          false,       //Clear all tag options
		transform:            "snakecase", //Transform adds a transform rule when adding tags. Current options: [snakecase, camelcase, lispcase, pascalcase, keep]
		sort:                 false,       //Sort sorts the tags in increasing order according to the key name
		override:             false,       //Override current tags when adding tags
		skipUnexportedFields: false,       //Skip unexported fields
	}
}

func OverrideJSON_File(fileName string, structName string, fieldName string, jsonName string, withOmitEmpty bool) error {
	cfg := newDefaultConfig()
	cfg.file = fileName
	cfg.structName = structName
	cfg.fieldName = fieldName
	cfg.write = true
	cfg.override = true
	cfg.skipUnexportedFields = true

	var flagModified bool
	var flagAddTags, flagAddOptions, flagRemoveTags, flagRemoveOptions string

	flagAddTags = "json:" + jsonName
	if withOmitEmpty {
		flagAddOptions = "json=omitempty"
	}
	/* override original function realMain */

	if flagModified {
		cfg.modified = os.Stdin
	}

	if flagAddTags != "" {
		cfg.add = strings.Split(flagAddTags, ",")
	}

	if flagAddOptions != "" {
		cfg.addOptions = strings.Split(flagAddOptions, ",")
	}

	if flagRemoveTags != "" {
		cfg.remove = strings.Split(flagRemoveTags, ",")
	}

	if flagRemoveOptions != "" {
		cfg.removeOptions = strings.Split(flagRemoveOptions, ",")
	}

	err := cfg.validate()
	if err != nil {
		return err
	}

	node, err := cfg.parse()
	if err != nil {
		return err
	}

	start, end, err := cfg.findSelection(node)
	if err != nil {
		return err
	}

	rewrittenNode, errs := cfg.rewrite(node, start, end)
	if errs != nil {
		if _, ok := errs.(*rewriteErrors); !ok {
			return errs
		}
	}

	out, err := cfg.format(rewrittenNode, errs)
	if err != nil {
		return err
	}

	if !cfg.write {
		fmt.Println(out)
	}
	return nil
}

func AddValidate_File(fileName string, structName string, fieldName string, validationRuleTag []string) error {

	cfg := newDefaultConfig()
	cfg.file = fileName
	cfg.structName = structName
	cfg.fieldName = fieldName
	cfg.write = true

	var flagModified bool
	var flagAddTags, flagAddOptions, flagRemoveTags, flagRemoveOptions string
	flagAddTags = "validate:" + strings.Join(validationRuleTag, "|")

	/* override original function realMain */

	if flagModified {
		cfg.modified = os.Stdin
	}

	if flagAddTags != "" {
		cfg.add = strings.Split(flagAddTags, ",")
	}

	if flagAddOptions != "" {
		cfg.addOptions = strings.Split(flagAddOptions, ",")
	}

	if flagRemoveTags != "" {
		cfg.remove = strings.Split(flagRemoveTags, ",")
	}

	if flagRemoveOptions != "" {
		cfg.removeOptions = strings.Split(flagRemoveOptions, ",")
	}

	err := cfg.validate()
	if err != nil {
		return err
	}

	node, err := cfg.parse()
	if err != nil {
		return err
	}

	start, end, err := cfg.findSelection(node)
	if err != nil {
		return err
	}

	rewrittenNode, errs := cfg.rewrite(node, start, end)
	if errs != nil {
		if _, ok := errs.(*rewriteErrors); !ok {
			return errs
		}
	}

	out, err := cfg.format(rewrittenNode, errs)
	if err != nil {
		return err
	}

	if !cfg.write {
		fmt.Println(out)
	}
	return nil
}

func AddDefault_File(fileName string, structName string, fieldName string, defaultValue string) error {

	cfg := newDefaultConfig()
	cfg.file = fileName
	cfg.structName = structName
	cfg.fieldName = fieldName
	cfg.write = true

	var flagModified bool
	var flagAddTags, flagAddOptions, flagRemoveTags, flagRemoveOptions string
	flagAddTags = "default:" + defaultValue

	/* override original function realMain */

	if flagModified {
		cfg.modified = os.Stdin
	}

	if flagAddTags != "" {
		cfg.add = strings.Split(flagAddTags, ",")
	}

	if flagAddOptions != "" {
		cfg.addOptions = strings.Split(flagAddOptions, ",")
	}

	if flagRemoveTags != "" {
		cfg.remove = strings.Split(flagRemoveTags, ",")
	}

	if flagRemoveOptions != "" {
		cfg.removeOptions = strings.Split(flagRemoveOptions, ",")
	}

	err := cfg.validate()
	if err != nil {
		return err
	}

	node, err := cfg.parse()
	if err != nil {
		return err
	}

	start, end, err := cfg.findSelection(node)
	if err != nil {
		return err
	}

	rewrittenNode, errs := cfg.rewrite(node, start, end)
	if errs != nil {
		if _, ok := errs.(*rewriteErrors); !ok {
			return errs
		}
	}

	out, err := cfg.format(rewrittenNode, errs)
	if err != nil {
		return err
	}

	if !cfg.write {
		fmt.Println(out)
	}
	return nil
}
