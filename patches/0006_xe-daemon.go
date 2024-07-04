diff --git a/xe-daemon/xe-daemon.go b/xe-daemon/xe-daemon.go
index 7b07873..70f474a 100644
--- a/xe-daemon/xe-daemon.go
+++ b/xe-daemon/xe-daemon.go
@@ -1,10 +1,10 @@
 package main
 
 import (
-       guestmetric "../guestmetric"
-       syslog "../syslog"
-       system "../system"
-       xenstoreclient "../xenstoreclient"
+       guestmetric "vyos-xe-guest-utilities/guestmetric"
+       syslog "vyos-xe-guest-utilities/syslog"
+       system "vyos-xe-guest-utilities/system"
+       xenstoreclient "vyos-xe-guest-utilities/xenstoreclient"
        "flag"
        "fmt"
        "io"
