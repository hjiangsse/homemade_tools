# homemade_tools
some homemade tools to increase working efficiency

# tools list
## 1. owl
fast doing ssh password-ness bettween two mechines, usage:  
```
owl -u [user] -h [hostaddr]
```  
example:  
```
own -u jingle -h 192.168.2.129
```

## 2. persona
do string(pattern) repalcement in all regular writeable files under current directory, usage:  
```
persona -o oldstring -n newstring  
```
example:
``` 
persona -o shanghai -n zhangjiang
```

## 3. newborn
do file(s) rename in current directory (example: test.txt --> test.bak), usage:
```  
newborn name --from exe --to txt
```  
the privious command will change "exe" to "txt" in all file name in current directory  
```
newborn content --from jingle --to mingle  
```
the privious command will change "exe" to "txt" in all file content in current directory

## 4. gitacc
accelerate git add, git commit, and git push process
``` bash
NAME:
   gitacc.exe - A new cli application

USAGE:
   gitacc.exe [global options] command [command options] [arguments...]

COMMANDS:
   fastpush, fp  fast push
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```