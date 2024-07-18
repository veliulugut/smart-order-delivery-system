import os 
from functools import lru_cache
from pydantic import BaseSettings
from pathlib import Path
from dotenv import load_dotenv
from urllib.parse import quote_plus

env_path = Path('.') / '.env'
load_dotenv(dotenv_path=env_path)

#App
APP_NAME: str = os.getenv("APP_NAME","FastAPI")


#Database Configuration
POSTGRES_HOST: str = os.getenv("POSTGRES_HOST","localhost")
POSTGRES_USER: str = os.getenv("POSTGRES_USER","root")
POSTGRES_PASSWORD: str = os.getenv("POSTGRES_PASSWORD","secret")
POSTGRES_PORT: str = os.getenv("POSTGRES_PORT","5432")
POSTGRES_DB: str = os.getenv("POSTGRES_DB","user")
DATABASE_URL = f"postgresql://{POSTGRES_USER}:{POSTGRES_PASSWORD}@{POSTGRES_HOST}:{POSTGRES_PORT}/{POSTGRES_DB}"

JWT_SECRET: str = os.getenv("JWT_SECRET","649fb93ef34e4fdf4187709c84d643dd61ce730d91856418fdcf563f895ea40f")
JWT_ALGORITHM: str = os.getenv("JWT_ALGORITHM","HS256")
ACCESS_TOKEN_EXPIRE_MINUTES: int = int(os.getenv("ACCESS_TOKEN_EXPIRE_MINUTES",3))
REFRESH_TOKEN_EXPIRE_MINUTES: int = int(os.getenv("REFRESH_TOKEN_EXPIRE_MINUTES",1440))


#APP Secret Key
SECRET_KEY:str= os.environ.get("SECRET_KEY","8deadce9449770680910741063cd0a3fe0acb62a8978661f421bbcbb66dc41f1")


@lru_cache()
def get_settings() -> Settings:
    return Settings()
