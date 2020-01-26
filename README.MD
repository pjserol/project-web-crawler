### Web Crawler

# Step 1 - call and get the content from the web page
# Step 2 - parse the document to find the title and the links
# Step 3 - create the node to store the result

### Run the project

# First param website
# Second param depth

* go run . https://godoc.org/ 2

### Test locally

* go test -v 

### Improvement

# Add test
# Remove links already visited
# Check only links who match the base
# Add concurrency
# Manage the case where the depth is negative to find the depth of the web site