
# This fetches the json and get the superhero that his id is 70 from those string name, wrapped in quotes the forward slash  is just a literal quote
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq -r '.[] | select(.id==70) | "\"\(.name)\""'

