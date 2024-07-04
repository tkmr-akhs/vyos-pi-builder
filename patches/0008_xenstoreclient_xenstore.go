diff --git a/xenstoreclient/xenstore.go b/xenstoreclient/xenstore.go
index 571481a..437537f 100644
--- a/xenstoreclient/xenstore.go
+++ b/xenstoreclient/xenstore.go
@@ -1,7 +1,7 @@
 package xenstoreclient
 
 import (
-       syslog "../syslog"
+       syslog "vyos-xe-guest-utilities/syslog"
        "bufio"
        "bytes"
        "encoding/binary"
