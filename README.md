# BRI RECE APP

## Create Database Name
#### ```DB_Rece```
### Create Extension Database PostgreSQL
##### ```CREATE EXTENSION IF NOT EXISTS "uuid-ossp";```

## Running App
#### ```go run api/main/main.go```

## Documentation API
#### USER
##### - Create User (POST)
###### ```localhost:8081/users```
![Screenshot from 2021-09-20 10-55-31](https://user-images.githubusercontent.com/62390363/133956675-14989391-ce15-4db3-9f13-3fbe44491cc0.png)

##### - Update User (PUT)


#### ACCOUNT
##### - Create Account (POST)
###### ```localhost:8081/account```
![Screenshot from 2021-09-20 10-55-49](https://user-images.githubusercontent.com/62390363/133956810-5e7fc444-b269-4300-bbfa-f3617c6e5948.png)

##### - In-active Account (PUT)
###### ```localhost:8081/account/inactive```


##### - Forgot Password (PUT)
###### ```localhost:8081/account/forgot```


##### - Login (POST)
###### ```localhost:8081/account/login```
![Screenshot from 2021-09-21 11-06-10](https://user-images.githubusercontent.com/62390363/134135283-b26ff99d-db86-460f-ab94-dd923327a29c.png)

##### - Edit Account (PUT)
###### ```localhost:8081/account```


#### WALLET
##### - WithDraw (PUT)
####### ```localhost:8081/wallet/withdraw```
###### Success
![Screenshot from 2021-09-21 15-16-52](https://user-images.githubusercontent.com/62390363/134136483-3e81d0e2-b713-47a6-8a3e-af58dd038ae3.png)
###### very low balance


##### - Top Up (PUT)
###### ```localhost:8081/wallet/topup```
![Screenshot from 2021-09-21 15-14-52](https://user-images.githubusercontent.com/62390363/134136211-d59cd1f0-f20b-4ffb-98a2-cee3059d808a.png)

##### - Get Wallet By Id (GET)
###### ```localhost:8081/wallet/{id}```
![Screenshot from 2021-09-21 15-06-38](https://user-images.githubusercontent.com/62390363/134135573-8a9b925a-0c59-4187-ab50-9dcd5d003c67.png)
![Screenshot from 2021-09-21 15-06-47](https://user-images.githubusercontent.com/62390363/134135870-5e08e842-5674-4385-b76a-5fc9cff3baca.png)
![Screenshot from 2021-09-21 15-06-58](https://user-images.githubusercontent.com/62390363/134135954-4f818727-22c2-4373-9125-6cb76eb9a8a4.png)

#### WALLET HISTORY
##### - GET TRANSACTION HISTORY (GET)
###### ```localhost:8081/history```