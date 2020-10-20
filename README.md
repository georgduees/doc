# doc
Days Of Code Helper, This Tool creates Logs and Helps with Docs
It can also create shareable links for Twitter.

It takes commit messages and generates Daily Log Markdown files.

# Features

- [ ] Parse commandline arguments
- [ ] Help Command with overview

- [ ] Git Integration
- [ ] Explore needed Git commands
    - [ ] Get current Git Repository name
    - [ ] Get Submodules as List
    - [ ] Create Submodule and Register
    - [ ] Parse Markdown Logfile
    - [ ] Load Markdown Template
    - [ ] Generate Markdown Logfile
    - [ ] Append to Logfile on commands
    - [ ] Commit changes to logfile to Git
    - [ ] Add Commits in module to "Main"-Project

# Information
DOC - abbreviation for DaysOfCode, or Doc(Document things...)
Functions with short description
DOC init
    Start with the program. Sets up basic structure doc.json as configuration file and adds it to the git repo.
DOC project list
    Show a List of current existing Projects assigned to the Repo and States(Last commit/ worked on)
DOC project begin PROJECTNAME
    Creates a Git Submodule with the name PROJECT and makes the Tool aware of it 
    add to Config.json
DOC project end PROJECTNAME
    Finishes the Project and creates final Logentry(Project summary)
DOC project log 
    Helps create Log entry for current Project
DOC project change
    Switches the current Working Project
DOC project share
    Creates Shareable information of latest activites/day/project

DOC day begin
    Starts timer to work on selected Project
DOC day end
    Create log file in master of 100days with summary what has been worked on since day begin.
DOC day log
    creates log entry for day
DOC day change
    Switch to different day/make changes to current day if already commits
DOC day share
    Creates shareable link information for current day


# ToDo:

# Upcoming Features


