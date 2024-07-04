diff --git a/xenstore/xenstore.go b/xenstore/xenstore.go
index 912de25..5ac272f 100644
--- a/xenstore/xenstore.go
+++ b/xenstore/xenstore.go
@@ -1,7 +1,7 @@
 package main
 
 import (
-       xenstoreclient "../xenstoreclient"
+       xenstoreclient "vyos-xe-guest-utilities/xenstoreclient"
        "fmt"
        "os"
        "strings"
