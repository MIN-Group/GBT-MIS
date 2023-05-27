# GBT-MIS Blockchain
## Run

  -  Quick start

  ```shell script
  go build -o GBT-MIS
  ./GBT-MIS
  ```
  >This method is to start 4 blockchain nodes locally

   -  General form

  ```shell script
  ./GBT-MIS -f <config_file>
  ```

  >You can modify the configuration file according to the actual situation or use the configuration file generation tool under common/config.go to generate a new configuration file. If it fails to run, please check whether the MongoDB database service is enabled. If it is not enabled, you can use the following command to enable it:                                                                                                                                      

  ```shell script
  sudo service mongod start
  ```

  >If MongoDB has not been installed, please refer to the following webpage to install it:                                                                                                                                       

  <https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/>


## Installation
  - Go mod environment
  ```shell script
   1. Set GOMODULE on
     		go env -w GO111MODULE="on"
   2. Update code
      After pulling the latest code, Goland will prompt you to detect Go Moudule. You need to click Enabled. Note: The following environment can be filled in or not. If necessary, fill in GOPROXY="https://proxy.golang.com.cn"
   3. Change Go Proxy
     		go env -w GOPROXY="https://proxy.golang.com.cn"
      Then type in
     		go env
      To check status
   4. Update go mod
     		go mod tidy
  ```

  -  Go dependencies

  ```shell script
    # go mod will automatically detect mod files and install dependencies
    go get github.com/Hyperledger-TWGC/ccs-gm
    go get github.com/JodeZer/mgop 
    go get github.com/google/uuid 
    go get github.com/karlseguin/ccache/v2 
    go get github.com/larspensjo/config 
    go get github.com/tinylib/msgp 
    go get github.com/yudeguang/ratelimit 
    go get gopkg.in/alexcesaro/quotedprintable.v3 
    go get gopkg.in/check.v1 
    go get gopkg.in/gomail.v2 
    go get gopkg.in/mgo.v2 
    go get gopkg.in/yaml.v2
    ...... 
  ```

## Instructions for use

  By default (quick start), four blockchain nodes are running locally, with an IP address of 127.0.0.1, blockchain communication ports `5010, 5011, 5012, 5013`, and service ports `8010, 8011, 8012, 8013` for the management front-end.
  The connection establishment ports provided for vpn-management-server are `9999, 10000, 10001, 10002`, and the ports for normal communication are `6666, 6667, 6668, 6669`.

- Module division

  ```textmate
    AccountManager		|	Account management module
       Message			|	Message management module
       MetaData			|	Definition of metadata format
       Database			|	Database module
       Network			|	Network module
      Blockchain			|	Program core module
    TransactionPool		|	Transaction pool module
       security			|	Security module
       common				|	Configuration file management module
       utils		    	|	Toolkit module
  ```

-  Configuration file description

  ```textmate
  [Log]
  "LogToFile"                 Whether to output the log to a file   
  "LogPath"                   Log file path
  "Level"                     Log display level: Panic 0,Fatal 1,Error 2,Warn 3,Info 4,Debug 5,Trace 6
  
  [Node]
  "WorkerList"                IP and port of bookkeepers
  "WorkerCandidateList"       IP and port of candidate bookkeepers
  "VoterList"                 IP and port of voters
  "BcManagementServerList"    Servers of blockchain management
  "ServerNum"	              	Servers number of blockchain management
  "SingleServerNodeNum"       Number of nodes running on this machine
  "IP"                        Local IP address
  "PubIP"                     Local public network address (Set to empty if none)
  "Port"                      Local port
  "Hostname"		          	Local node name
  "Areaname"		          	Name of the region where the node is located
  "Countryname"		        Name of the country where the node is located
  "Longitude"                 Longitude of the node's location
  "Latitude"                  Latitude of the node's location
  "CacheTime"                 Duration of caching node status information, in minutes
  "IsNewJoin"                 Indicates whether it is a newly joined node
  
  [Consensus] 
  "PubkeyList"                Public key of the node running on this machine
  "PrikeyList"                Private key of the node running on this machine
  "MyPubkey"                  Reserved configuration
  "MyPrikey"                  Reserved configuration
  "GenesisDutyWorker"         Node number that generated the genesis block
  "WorkerNum"                 Number of bookkeepers
  "VotedNum"                  Number of voters
  "BlockGroupPerCycle"        Rotation cycle of bookkeepers
  "Tcut"                      Overtime time
  "GenerateBlockPeriod"       Length of Generate block cycle
  "TxPoolSize"                Size of transaction pool
  
  [MIR]
  "SqlitePath"                File path sqlite database
  
  [SESSION]
  "DefaultExpiration"         Default validity period of the session, in minutes
  "CleanupInterval"           Cleaning cycle of a session, in minutes
  ```