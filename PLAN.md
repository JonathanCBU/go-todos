## Phase 1: Define Your Data Model

Start with `internal/models/record.go` - define what a todo item looks like:

- ID, Title, Description, Status (pending/completed), Created/Updated timestamps, priority

### Future Feature

- Due dates, categories/tags

## Phase 2: Core CSV Operations

Build your CSV handlers in `internal/csv/`:

- **reader.go**: Load todos from CSV file, handle missing file gracefully
- **writer.go**: Save todos to CSV file, backup existing file
- Choose a standard CSV format (headers, field order, how to handle commas in text)

## Phase 3: Basic CLI Commands

Implement in `cmd/` using Cobra framework:

- **root.go**: Set up base command, global flags (like CSV file path)
- **add.go**: Add new todo items
- **list.go**: Display todos (all, by status, with filtering)
- **complete.go**: Mark todos as done
- **remove.go**: Delete todos

## Phase 4: Essential Features

- Auto-generate unique IDs for new todos
- Input validation (required fields, reasonable limits)
- Pretty-print output with tables or formatting
- Handle edge cases (empty file, corrupted CSV, file permissions)

## Phase 5: Quality of Life Improvements

- Search/filter todos by text, status, date ranges
- Edit existing todos
- Different output formats (JSON, simple text)
- Configuration file for default settings

## Phase 6: Polish

- Comprehensive error handling and user-friendly messages
- Unit tests for your CSV operations
- Documentation and help text
- Consider data persistence location (home directory, configurable path)

**Start with Phase 1-2 to get your data foundation solid, then build the CLI interface on top. Want me to dive deeper into any specific phase?**
