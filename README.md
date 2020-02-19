[![Go Report Card](https://goreportcard.com/badge/github.com/LucHighwalker/nor)](https://goreportcard.com/report/github.com/LucHighwalker/nor)

# JAMF - Just Another Mvc Framework

Jamf is just another mvc framework. Spawned by me wanting to automate a lot of the repetetive setup that I go through when creating projects. It's currently still very much a work in progress and thus very limited. 

| Some limitations: |
| --- |
|   - Only supports NodeJs (typescript) |
|   - Only supports MongoDb |
|   - Model Controller integration must be done manually |
|   - Views are not yet implemented |

# Usage 

## Initialize a project

```
jamf init [project-name]
```

| Flags: |
| --- |
| defPort - The default port that the app will run on. |
| dbPort - The default port that the database uses locally. |

## Generate Model

```
jamf model [model-name] [field-name]:[field-type]:[field-options]
```

| Field Options: |
| --- |
| required - Sets required to true. |
| min=[int] - The minumum amount of characters this field requires. |
| max=[int] - The maximum amount of characters this field allows. |
| validate{[regexp]} - Compares the field to the input regular expression. |

## Generate Controller
nor c SomeController test:public:a=string:b=number another:private:a=number:b=date

```
jamf controller [controller-name] [action-name]:[public || private]:[action-parameter]=[paramter-type]
```

