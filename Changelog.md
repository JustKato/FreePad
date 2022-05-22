# 1.1.0
Implemented a rate-limiting system, quite primitive and basic implementation on my part since it's looking at all requests not just the POST requests, this can be bad news if someone is using the service a ton and won't truly protect from floods as it's ip-based but should offer a level of security better than none.

# 1.0.0
Initial release of FreePad, this included all of the basic functionality such as:
- Homepage
- Generating Pads
- Real-Time Saving
- Download Functionality
- Archive Functionality
- Refresh Button
- Dark/Light Theme Toggles
- New look/logo

# 0.9
Old version of FreePad which depended on a database for storing data, this was later dropped as it was pointless since we are only storing temporary data, so we moved to "v2"