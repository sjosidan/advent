# importing the requests library
import requests


#SBOX
# api-endpoint
URL = "https://api.sandbox.hierarchy.enlight.skf.com/nodes/2b3d630b-6f87-4377-bbb1-cb9cfb352fa0/subtree?type=asset"

COMPURL1 = "https://api.sandbox.hierarchy.enlight.skf.com/assets/"
COMPURL2 = "/components?limit=200&offset=0"


UPDATEURL1 = "https://api.sandbox.hierarchy.enlight.skf.com/assets/"
UPDATEURL2 = "/components/"

#PROD
#URL = "https://api.hierarchy.enlight.skf.com/nodes/4d7c1138-c973-48ce-95fe-48af50ecd61e/subtree?type=asset"


#COMPURL1 = "https://api.hierarchy.enlight.skf.com/assets/"

#REPLACE BEARING
replaceME = "APANS"

replaceMAN = "SKF"
replaceDES = "6312"



# location given here
keyy = ""
#keyy = "" 
# defining a params dict for the parameters to be sent to the API
PARAMS = {'authorization':keyy}
  
r = requests.get(url = URL, headers = PARAMS, timeout=2)
calls = 1 
# extracting data in json format
data = r.json()
r.close()
#print(data)  
assets = data['nodes']
print("\n\n\n\n\n")
res = []
for a in assets : 
      url2 = COMPURL1 + a['id'] + COMPURL2
      calls = calls + 1
      #r2 = requests.get(url = url2, headers = PARAMS, timeout=4 )
      try:
            r2 = requests.get(url = url2, headers = PARAMS, timeout=(2,4))
            #print(response.status_code)
            data2 = r2.json()
            r2.close()      

            comps = data2['components']
            if len(comps) > 0 :
                  for c in comps : 
                        bearings = c['type']
                        if c['type'] == "bearing" :
                              if c['manufacturer'] == "" :
                                    print("FIX", c['designation'])
                                    res.append(c['designation'])
                                    url3 = UPDATEURL1 + a['id'] + UPDATEURL2 + c['id']
                                    c['designation'] = replaceDES
                                    c['manufacturer'] = replaceMAN
                                    print("UPDATE", url3, c)
                                    try: 
                                          r3 = requests.put(url3, headers=PARAMS, data=c, timeout=(2,4))
                                          print(r3)
                                          r3.close()
                                    except requests.exceptions.RequestException as err :
                                          print(err)
      except requests.exceptions.RequestException as e:
            print(e)
      print("Call: ", calls)

uniqueList = list(set(res))
#print(uniqueList)
for r in uniqueList : 
      print(r, "\t\t", res.count(r))
