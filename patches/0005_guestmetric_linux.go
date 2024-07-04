diff --git a/guestmetric/guestmetric_linux.go b/guestmetric/guestmetric_linux.go
index e61ccca..653a86b 100644
--- a/guestmetric/guestmetric_linux.go
+++ b/guestmetric/guestmetric_linux.go
@@ -1,7 +1,7 @@
 package guestmetric
 
 import (
-       xenstoreclient "../xenstoreclient"
+       xenstoreclient "vyos-xe-guest-utilities/xenstoreclient"
        "bufio"
        "bytes"
        "fmt"
