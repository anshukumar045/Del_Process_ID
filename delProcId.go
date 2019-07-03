package utility
import (
        "io/ioutil"
        "strconv"
        "bufio"
        "os"
        "strings"
        "path/filepath"
        "log"
        "time"
        "os/exec"
)

// remove duplicates from slice
func unique(intSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{}
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
return list
}

// Get the files in the directory
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	 if !info.IsDir() {
	  files = append(files, path)
	 }
	 return nil
	})
return files, err
}
   
// Get the file modification time
func getFileModTime(fname string) time.Time {
	filestat, err := os.Stat(fname)
	if err != nil {
		log.Println("ERROR to find creation time:- ",err, " file :- ", fname)
	}
return filestat.ModTime()
}

// Find the PID from and parse the command line for the partion
func getCommandline(pid int)(string, error){
	path := "/proc/"+strconv.Itoa(pid)+"/cmdline"
	cln, err :=  exec.Command("cat", path).Output()
	if err != nil {
		return " ", err
	}
return string(cln), err
}

// Find the PID and kill the process
func getPIDandDel(part string) {
	filelst, err := ioutil.ReadDir("/proc")
	if err !=nil {
			log.Println(err)
	}
	for _, f := range filelst {
		pid, err := strconv.Atoi(f.Name())
		if err == nil {
			if checkUser("/proc/"+f.Name()) == true {
				cmdline , err := getCommandline(pid)   // command line for each PID
				if err != nil {
						log.Println(err)
				}
				if strings.Contains(cmdline, part) {
						cmd := exec.Command("kill", "-9",f.Name())
						err := cmd.Run()
						if err != nil {
								log.Println("pid= ",pid)
								log.Printf(" getPIDandDel Pid delition finished with error %v", err)
						} else {
						log.Println("pid = ", pid, " was killed")
						}
				}
			}
		}
	}
}

// CHECK FOR THE USER OF THE PROCESS
func checkUser(fname string) bool {
	username, err := exec.Command("stat","-c", "%U", fname).Output()
	if err != nil {
			log.Fatal("error in finding username ANSHU ", err)
	}
	if strings.TrimSpace(string(username)) == "anshu" {
			return true
	} else {
			return false
	}
}


