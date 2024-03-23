package helper

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"martpay/database"
	"strings"
)

func CustomRules() {
	// Exists validation to ensure the value of the field exist in the specified table for the column e.g. exists:table,column.
	// Column will be taken as the field if not specified in the rule
	govalidator.AddCustomRule("exists", func(field string, rule string, message string, value any) error {
		column := field
		values := strings.Split(strings.TrimPrefix(rule, "exists:"), ",")
		if len(values) == 0 {
			return fmt.Errorf("The table is required for the validation of %s", field)
		}

		table := values[0]

		if len(values) > 1 {
			column = values[1]
		}

		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE %s = ?", table, column)

		var count int

		database.Db.Raw(query, value).Scan(&count)

		if count < 1 {
			return fmt.Errorf("The %s doesn't exist in the %s table", column, table)
		}

		return nil
	})

	// Unique validation to ensure the value of the field doesn't exist in the specified table for the column e.g. unique:table,column.
	// Column will be taken as the field if not specified in the rule.
	govalidator.AddCustomRule("unique", func(field string, rule string, message string, value any) error {
		column := field
		values := strings.Split(strings.TrimPrefix(rule, "unique:"), ",")
		if len(values) == 0 {
			return fmt.Errorf("The table is required for the validation of %s", field)
		}

		table := values[0]

		if len(values) > 1 {
			column = values[1]
		}

		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE %s = ?", table, column)

		var count int

		database.Db.Raw(query, value).Scan(&count)

		if count > 0 {
			return fmt.Errorf("The %s already exists in the %s table", column, table)
		}

		return nil
	})

	// Confirm the field matches the specified value.
	govalidator.AddCustomRule("confirm", func(field string, rule string, message string, value any) error {
		originalValue := strings.TrimPrefix(rule, "confirm:")
		originalField := strings.TrimPrefix(field, "confirm_")

		if originalValue == rule {
			return fmt.Errorf("The value is required for confirmation")
		}

		if originalValue != value {
			return fmt.Errorf("The %s does not match", originalField)
		}

		return nil
	})
}
