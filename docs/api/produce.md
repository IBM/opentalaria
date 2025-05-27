# Produce API

# Caveats
OpenTalaria supports only Record sets with Magic Byte 2. The broker cannot handle the old message format and will throw an error if a producer tries to use it.

For now OpenTalaria does not support transactional producers. This will change in later releases.