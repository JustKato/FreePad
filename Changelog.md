# 1.3.0 👀
Implemented a views system, now everyone can see how many times a pad has been accessed, an auto-save has also been added for those views to file in the `data` dir.

# 1.2.0 🍥
QR Code generation has been implemented, the refresh button has been removed for the sake of keepign things simple and symmetrical. This will generate a QR code on the javascript side. 

I have created this feature mainly for people that are trying to quickly transfer information from their computer to their phone extremely quickly, simply just type your Pad, click on the `QR Code` button and then pull your phone's camera out and get instant access, no longer typing the URL or sending the URL through some messaging app and wastring time.

# 1.1.1 🗒
The freepad version has been added as a header to the response

# 1.1.0 🛑
Implemented a rate-limiting system, quite primitive and basic implementation on my part since it's looking at all requests not just the POST requests, this can be bad news if someone is using the service a ton and won't truly protect from floods as it's ip-based but should offer a level of security better than none.

# 1.0.0 🖥
Initial release of FreePad, this included all of the basic functionality such as:
- Homepage
- Generating Pads
- Real-Time Saving
- Download Functionality
- Archive Functionality
- Refresh Button
- Dark/Light Theme Toggles
- New look/logo

# 0.9 🥶
Old version of FreePad which depended on a database for storing data, this was later dropped as it was pointless since we are only storing temporary data, so we moved to "v2"