package main
import "fmt"
import "os"
import "io/ioutil"
import "encoding/json"
import "strconv"
import "strings"



type ClusterNodeInfo struct {
    URI              string `json:"uri"`
    Status           string `json:"status"`
    State            string `json:"state"`
    RegistrationDate string `json:"registrationDate"`
    Role             string `json:"role"`
}

type ClusterInfo struct {
    UUID             string `json:"uuid"`
    Name             string `json:"name"`
    Type             string `json:"type"`
    Version          string `json:"version"`
    RegistrationDate string `json:"registrationDate"`
    Status           string `json:"status"`
    State            string `json:"state"`
    Nodes            []ClusterNodeInfo `json:"nodes"`
}
// TODO: Nao tem uuid ? VF3 atual tem ...
//  "uuid":"39373638-3136-584d-5138-313430305934",
//  "uri":"https://vf-3.labs.hpecorp.net:20443/api/v1/devices/39373638-3136-584d-5138-313430305934",
//  "status":"critical",
//  "state":"registered",
//  "role":"worker",
//  "registrationDate":"Fri, 15 May 2020 16:12:51 UTC"}

type ClustersInfo struct {
	Members []ClusterInfo `json:"members"`
}

type SecurityCentralInfo struct {
    clusterName      string
    clusterType      string
    NodeName         string
    NodeProductName  string
    ProfileName      string
    WorkloadProfileName string
}


func GetSecurityCentralInfo(targetUuid string) (SecurityCentralInfo, error) {
//    cluster := clusters.Members[0]
    var info SecurityCentralInfo
    info.clusterName = "not found"
    info.clusterType = "not found"
    info.NodeName = "not found"
    info.NodeProductName = "not found"
    info.ProfileName = "not found"
    info.WorkloadProfileName = "not found"

    // substituir
    clusters := GetClustersInfo()
    info.clusterName,info.clusterType = getClusterNameType(targetUuid, clusters)
    return info,nil
}



func GetClusterInfo() (ClusterInfo, error) {
    clusters := GetClustersInfo()
    cluster := clusters.Members[0]
    return cluster,nil
}


func GetClustersInfo() (ClustersInfo) {

   // Open our jsonFile
    jsonFile, err := os.Open("clusters-body.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
        // TODO: how to exit ?
    }
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened jsonFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users array
    var clusters ClustersInfo

    // we unmarshal our byteArray which contains our jsonFile's content into 'users' which we defined above
    json.Unmarshal(byteValue, &clusters)

    fmt.Println("Status: " + clusters.Members[0].Name)
    return clusters
}

func main() {
    targetUuid := "325a2f60-d854-47ae-b760-b08341816aeb" // "cdf356f1-8e3a-6bc9-04d7-a24f738d665f"
    clusters := GetClustersInfo()

    fmt.Println("Status: " + clusters.Members[0].Name)
    
    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
    for i := 0; i < len(clusters.Members); i++ {
        fmt.Println()
        fmt.Println("============ CLUSTER " + strconv.Itoa(i))
        fmt.Println("uuid : " + clusters.Members[i].UUID)
        fmt.Println("name : " + clusters.Members[i].Name)
        fmt.Println("type : " + clusters.Members[i].Type)
        fmt.Println("Version : " + clusters.Members[i].Version)
        fmt.Println("RegistrationDate : " + clusters.Members[i].RegistrationDate)
        fmt.Println("Status : " + clusters.Members[i].Status)
        fmt.Println("State : " + clusters.Members[i].State)
        for j:= 0; j<len(clusters.Members[i].Nodes); j++ {
            //fmt.Println()
            fmt.Println("   Node " + strconv.Itoa(j))
            fmt.Println("   URI : " + clusters.Members[i].Nodes[j].URI)
            fmt.Println("   Status : " + clusters.Members[i].Nodes[j].Status)
            fmt.Println("   State : " + clusters.Members[i].Nodes[j].State)
            fmt.Println("   RegistrationDate : " + clusters.Members[i].Nodes[j].RegistrationDate)
            fmt.Println("   Role : " + clusters.Members[i].Nodes[j].Role)
        }
    }    

    fmt.Println("hello world 888")
    getClusterNameType(targetUuid, clusters)
}

func getClusterNameType(targetUuid string, clusters ClustersInfo) (string, string) {
    for i := 0; i < len(clusters.Members); i++ {
        for j:= 0; j<len(clusters.Members[i].Nodes); j++ {
            if strings.Contains(clusters.Members[i].Nodes[j].URI, targetUuid) {
                clusterName := clusters.Members[i].Name
                clusterType := clusters.Members[i].Type
                fmt.Println("=== ACHOU " + clusters.Members[i].Nodes[j].URI + " " + clusterName + " " + clusterType)
                return clusterName, clusterType
            }
        }
    }
    return "not found", "not found"
}