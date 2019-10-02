# gig
CLI for creating .gitignore files (based on gitignore.io)

## Disclaimer
All credits goes to gitignore.io for the actual generation of .gitignore files. This tools only handles programmatic validation and orchestration of getting the gitignore files. It only works as a terminal based wrapper for an API, where i had no involvement in the API parts.

Secondly, this is nothing special (more a teaching project for me where it was easier to have a set goal to work against). The functionality could just be to have an alias which downloads the gitignore file instead.

# Usage
* gig configure (fetches the templates for each languages that is available to use)
* gig list (lists the available languages found in the json file from the "configure" command)
* gig generate (generate the gitignore file by calling the gitignore.io API based on the parameters passed to this tool)
