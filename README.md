# Working with External APIs (Twitter API)
## How to interact with external APIs by working with the Twitter API and generating our own Twitter API keys

### Part 1 : Setup Twitter Developer Account and Generate API Keys

#### 1. Create a new Twitter account

##### 
- Click on the given [Link](https://x.com/i/flow/login) to signup if you don't have one
#### 2. Sign up for a twitter developer account

##### 
- Navigate to the Twitter Developer Platform. Click [here](https://developer.x.com/en)
- Sign up for a Developer account
- Create new project and app
  
#### 3. Generate API keys

##### 
- Click on Projects & Apps under Dashboard and now tap on Apps and then navigate to keys and tokens
- Generate the following keys & tokens:
  
    - API Key and Secret
    - Bearer Token
    - Access Token and Secret

![image](https://github.com/user-attachments/assets/28ae1d77-a249-4bed-b207-55f5e94a5d68)

  
#### 4. Configure OAuth

#####
- Click on settings, tap on set up button for user authentication

![image](https://github.com/user-attachments/assets/3354871c-357e-4a2c-88e4-ca39631389c3)

- Setting up of callback url

![image](https://github.com/user-attachments/assets/23ee3925-93b9-4748-b3fa-705158a37f2c)


### Part 2: Interact with the Twitter API

##### 
- Posting a new tweet
  
   - Let's post a tweet using POST statuses/update endpoint of the Twitter API
   - We can use go.

![image](https://github.com/user-attachments/assets/8fb72b9d-d364-4010-8771-52550b01f8f5)

   - We can see the tweet posted on the timeline

![image](https://github.com/user-attachments/assets/e86335c1-6471-43fe-8c89-7db94fb73e8a)

   
#####
- Deleting a tweet
  
   - Now using go language, deleting a tweet can be possible by passing the tweet ID. Let me delete the tweet that was posted initially and Tweet got deleted!


![image](https://github.com/user-attachments/assets/8ec8a42e-8adb-49a5-aa78-535feaae9ab7)






