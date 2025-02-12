                    App crud cơ bản có phân quyền, jwt token. 

B1. pull về

B2. tạo file .env có nội dung:


POSTGRES_USER=postgres

POSTGRES_PASSWORD=12345

POSTGRES_DB=postgres

DB_HOST=database

DB_PORT=5432

REFRESH_TOKEN_EXPIRATION=1440m  # 2 ngày

REFRESH_TOKEN_MAXAGE=1440     

REFRESH_TOKEN_SECRET=refresh_secret

ACCESS_TOKEN_EXPIRATION=5m      # 5 phút

ACCESS_TOKEN_SECRET=access_secret

REDIS_HOST=redis

REDIS_PORT=6379

REDIS_DB=0


B3. chạy docker-compose up --build


B4. mở web enter link: http://localhost:8081/swagger/index.html 


Các api trong link chạy localhost với port 8081
