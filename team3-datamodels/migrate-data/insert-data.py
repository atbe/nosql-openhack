from pprint import pprint
import json
import random
import azure.cosmos.cosmos_client as cosmos_client
import azure.cosmos.documents as doc
from azure.cosmos import CosmosClient, PartitionKey, exceptions

config = {
    'ENDPOINT': 'https://cosmosdb2-team3.documents.azure.com:443/',
    'PRIMARYKEY': 'uDw5eCXYABEpiz16QTKvd07GhCt68D9g9pNye6shFbalntrmHS1gI9NXCY0m8TOHjGYwtqjY7X2j19mq3Wt7hQ==',
    'DATABASE': 'AbeTemp',
    'CONTAINER': 'Cart'
}

# Initialize the Cosmos client
connection_policy = doc.ConnectionPolicy()
# Disable in production
connection_policy.DisableSSLVerification = "true"

client = cosmos_client.CosmosClient(url=config['ENDPOINT'], credential=config['PRIMARYKEY'])

database_name = config['DATABASE']
database = client.get_database_client(database_name)

container_name = config['CONTAINER']
try:
    container = database.create_container(id=container_name, partition_key=PartitionKey(path="/UserId"))
except exceptions.CosmosResourceExistsError:
    container = database.get_container_client(container_name)
except exceptions.CosmosHttpResponseError:
    raise

container = database.get_container_client(container_name)

# Iterating over fake person data and storing in DB
with open('carts.json', 'r') as infile:
	jsonData = json.load(infile)
	for cart in jsonData:
		del cart['Id']
		pprint(cart)
		container.upsert_item(cart)
