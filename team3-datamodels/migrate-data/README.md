# How to use

```bash
$ python3 -m venv .env
$ source .env/bin/activate
$ pip install -r requirements.txt
$ ./run.sh
```

The script will:

1. Assemble a file named `carts.json` based on the data in the sql database
2. Insert all objects in the `carts.json` into the Cosmos database