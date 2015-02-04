package junos

// Variables we reference, such as Content-Types, RPC commands.
var (
	contentDiscoverDevices = "application/vnd.net.juniper.space.device-management.discover-devices+xml;version=2;charset=UTF-8"
	contentExecDeploy      = "application/vnd.net.juniper.space.software-management.exec-deploy+xml;version=1;charset=UTF-8"
	contentExecRemove      = "application/vnd.net.juniper.space.software-management.exec-remove+xml;version=1;charset=UTF-8"
	contentExecStage       = "application/vnd.net.juniper.space.software-management.exec-stage+xml;version=1;charset=UTF-8"
	contentAddress         = "application/vnd.juniper.sd.address-management.address+xml;version=1;charset=UTF-8"
	contentUpdateDevices   = "application/vnd.juniper.sd.device-management.update-devices+xml;version=1;charset=UTF-8"
	contentPublish         = "application/vnd.juniper.sd.fwpolicy-management.publish+xml;version=1;charset=UTF-8"
	contentAddressPatch    = "application/vnd.juniper.sd.address-management.address_patch+xml;version=1;charset=UTF-8"
	contentService         = "application/vnd.juniper.sd.service-management.service+xml;version=1;charset=UTF-8"
	contentServicePatch    = "application/vnd.juniper.sd.service-management.service_patch+xml;version=1;charset=UTF-8"
	rpcCommand             = "<rpc><command format=\"text\">%s</command></rpc>"
	rpcCommandXML          = "<rpc><command format=\"xml\">%s</command></rpc>"
	rpcCommit              = "<rpc><commit-configuration/></rpc>"
	rpcCommitAt            = "<rpc><commit-configuration><at-time>%s</at-time></commit-configuration></rpc>"
	rpcCommitCheck         = "<rpc><commit-configuration><check/></commit-configuration></rpc>"
	rpcCommitConfirm       = "<rpc><commit-configuration><confirmed/><confirm-timeout>%d</confirm-timeout></commit-configuration></rpc>"
	rpcFactsRE             = "<rpc><get-route-engine-information/></rpc>"
	rpcFactsChassis        = "<rpc><get-chassis-inventory/></rpc>"
	rpcConfigFileSet       = "<rpc><load-configuration action=\"set\" format=\"text\"><configuration-set>%s</configuration-set></load-configuration></rpc>"
	rpcConfigFileText      = "<rpc><load-configuration format=\"text\"><configuration-text>%s</configuration-text></load-configuration></rpc>"
	rpcConfigFileXML       = "<rpc><load-configuration format=\"xml\"><configuration>%s</configuration></load-configuration></rpc>"
	rpcConfigURLSet        = "<rpc><load-configuration action=\"set\" format=\"text\" url=\"%s\"/></rpc>"
	rpcConfigURLText       = "<rpc><load-configuration format=\"text\" url=\"%s\"/></rpc>"
	rpcConfigURLXML        = "<rpc><load-configuration format=\"xml\" url=\"%s\"/></rpc>"
	rpcGetRescue           = "<rpc><get-rescue-information><format>text</format></get-rescue-information></rpc>"
	rpcGetRollback         = "<rpc><get-rollback-information><rollback>%d</rollback><format>text</format></get-rollback-information></rpc>"
	rpcGetRollbackCompare  = "<rpc><get-rollback-information><rollback>0</rollback><compare>%d</compare><format>text</format></get-rollback-information></rpc>"
	rpcHardware            = "<rpc><get-chassis-inventory/></rpc>"
	rpcLock                = "<rpc><lock><target><candidate/></target></lock></rpc>"
	rpcRescueConfig        = "<rpc><load-configuration rescue=\"rescue\"/></rpc>"
	rpcRescueDelete        = "<rpc><request-delete-rescue-configuration/></rpc>"
	rpcRescueSave          = "<rpc><request-save-rescue-configuration/></rpc>"
	rpcRollbackConfig      = "<rpc><load-configuration rollback=\"%d\"/></rpc>"
	rpcRoute               = "<rpc><get-route-engine-information/></rpc>"
	rpcSoftware            = "<rpc><get-software-information/></rpc>"
	rpcUnlock              = "<rpc><unlock><target><candidate/></target></unlock></rpc>"
	rpcVersion             = "<rpc><get-software-information/></rpc>"
)
