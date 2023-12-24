### Installation
#### Prerequistites:
Golang

**Steps to install:**
- `git clone https://github.com/ppadti/Go-Projects.git`
- `cd Go-Projects`
  
**For CLIExample project:**

- `cd CLIExample`

- `go build`

- `./CLIExample`
  
**For FileSystem project:**
  
- `cd FileSystem`
  
- `go build`
  
- `./CLIExample`
  

## [CLIExample](https://github.com/ppadti/Go-Projects/tree/main/CLIExample)
This command-line interface (CLI) project processes user input to format and display specific details based on arguments provided. 
The first argument (mandatory) can be a name, company, or hometown. It's essential and must be present; otherwise, the application will prompt an error message.

### Features:

**Input Validation:** Mandatory first argument (name, company, or hometown) ensures valid input for further processing.

**Optional Arguments:**

**Case:** Accepts two optional arguments - 'case' and 'full'.

**Case Option:** Converts the first argument to upper or lower case based on user input.

**Full Option:** Displays the full name (if available) based on the first argument input.


## [FileSystem](https://github.com/ppadti/Go-Projects/tree/main/FileSystem)
This project simplifies file management by creating a text file if it doesn't exist. 
It allows users to store data as key-value pairs and retrieve specific values by entering the corresponding key via the command-line interface (CLI). 
### Features:
**File Creation**: Automatic creation of the designated text file if not found.

**Key-Value Storage**: Efficient storage mechanism allowing users to associate values with unique keys.

**CLI Interaction**: Command-line interface for adding key-value pairs and retrieving stored values based on user-input keys.

