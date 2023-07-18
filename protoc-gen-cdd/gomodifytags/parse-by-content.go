package gomodifytags

import (
	"os"
	"strings"
)

func OverrideJSON_Content(content string, structName string, fieldName string, jsonName string, withOmitEmpty bool) (string, error) {
	cfg := newDefaultConfig()
	cfg.file = content
	cfg.structName = structName
	cfg.fieldName = fieldName
	cfg.write = false
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
		return "", err
	}

	node, err := cfg.parseByContent(content)
	if err != nil {
		return "", err
	}

	start, end, err := cfg.findSelection(node)
	if err != nil {
		return "", err
	}

	rewrittenNode, errs := cfg.rewrite(node, start, end)
	if errs != nil {
		if _, ok := errs.(*rewriteErrors); !ok {
			return "", errs
		}
	}

	out, err := cfg.format(rewrittenNode, errs)
	if err != nil {
		return "", err
	}

	return out, nil
}

func AddDefault_Content(content string, structName string, fieldName string, defaultValue string) (string, error) {

	cfg := newDefaultConfig()
	cfg.file = content
	cfg.structName = structName
	cfg.fieldName = fieldName
	cfg.write = false

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
		return "", err
	}

	node, err := cfg.parseByContent(content)
	if err != nil {
		return "", err
	}

	start, end, err := cfg.findSelection(node)
	if err != nil {
		return "", err
	}

	rewrittenNode, errs := cfg.rewrite(node, start, end)
	if errs != nil {
		if _, ok := errs.(*rewriteErrors); !ok {
			return "", errs
		}
	}

	out, err := cfg.format(rewrittenNode, errs)
	if err != nil {
		return "", err
	}

	return out, nil
}

func AddValidate_Content(content string, structName string, fieldName string, validationRuleTag []string) (string, error) {

	cfg := newDefaultConfig()
	cfg.file = content
	cfg.structName = structName
	cfg.fieldName = fieldName
	cfg.write = false

	var flagModified bool
	var flagAddTags, flagAddOptions, flagRemoveTags, flagRemoveOptions string
	flagAddTags = "validate:" + strings.Join(validationRuleTag, "|")
	flagAddTags = strings.ReplaceAll(flagAddTags, ",", ";")
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
		return "", err
	}

	node, err := cfg.parseByContent(content)
	if err != nil {
		return "", err
	}

	start, end, err := cfg.findSelection(node)
	if err != nil {
		return "", err
	}

	rewrittenNode, errs := cfg.rewrite(node, start, end)
	if errs != nil {
		if _, ok := errs.(*rewriteErrors); !ok {
			return "", errs
		}
	}

	out, err := cfg.format(rewrittenNode, errs)
	if err != nil {
		return "", err
	}

	return out, nil
}
