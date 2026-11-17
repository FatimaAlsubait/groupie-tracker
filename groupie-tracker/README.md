# Groupie Trackers

A web platform written in Go that consumes an external RESTful API and visualizes musical artists, their members, creation years, first album releases, and concert relations.

## Features
- **Catalogue View**: Displays all bands in a responsive card grid with creation dates and member counts.
- **Band Profiles**: Shows band members and tour schedules cleanly mapped by location and dates.
- **Client-Server Search Action**: Filters bands by artist name, member names, creation years, and first album dates using native HTTP GET form requests without JavaScript.
- **Resilient & Crash-Proof**: Implements panic recovery middleware, connection timeouts, and buffered template execution with custom HTTP error pages (400, 404, 405, 500).

## Requirements
- Go 1.22 or higher
- Only Go standard library packages used

## Who worked on the project
- Fatima Alsubait (Leader)
- Sheikha Alajmi 
- Lujayn Saleh

## Running the Project
Make sure to run the project from the root directory:

```bash
go run ./cmd/server/main.go