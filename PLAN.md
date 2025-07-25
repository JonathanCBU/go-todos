## CLI structure

- **root.go**: Set up base command, global flags (like CSV file path)
- **add.go**: Add new todo items
- **list.go**: Display todos
- **start.go** Mark todo as in progress
- **complete.go**: Mark todos as done
- **remove.go**: Delete todos

## Improved basics

- Handle missing csv file elegantly
- Auto-generate IDs
- Input validation (required fields, reasonable limits)
- Pretty-print output with tables or formatting

- Handle edge cases (empty file, corrupted CSV, file permissions)

## Phase 5: Quality of Life Improvements

- Edit existing todos
- Different output formats (JSON, simple text)
- Due dates, categories/tags
- Print todos with filters for priority and/or status

## Phase 6: Polish

- Comprehensive error handling and user-friendly messages
- Unit tests for your CSV operations
- Documentation and help text
- Editable config file for users
