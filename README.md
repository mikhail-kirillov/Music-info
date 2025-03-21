# Music-info  

Effective Mobile | Online Song Library REST API Test Task  

## Description  

Music-info is a REST API server designed for managing song information. It allows storing, updating, deleting, and retrieving song details, as well as fetching lyrics from an external server.  

## Technologies Used  

- **Golang**: Programming language used to implement the server.  
- **GIN**: Web framework for building applications in Go.  
- **slog**: Logging package.  
- **viper**: Library for handling configurations.  
- **testify**: Testing framework.  
- **swaggo**: Tool for generating API documentation.  
- **PostgreSQL**: Relational database.  
- **gorm**: ORM library for Go.  
- **Docker & Docker Compose**: Tools for containerization and application orchestration.  

## API Functionality  

- **GET /songs**: Retrieve a list of songs.  
- **POST /songs**: Add a new song.  
- **PUT /songs/{id}**: Update song details by ID.  
- **DELETE /songs/{id}**: Delete a song by ID.  
- **GET /songs/{id}/lyrics**: Fetch lyrics for a song by ID.  

## External API for Fetching Song Information  

To retrieve complete song details, the server interacts with an external API using the following parameters:  

- **GET /info**:  
  - Query parameters:  
    - `group` (string, required): Name of the band.  
    - `song` (string, required): Title of the song.  
  - Response:  
    - **200 OK**: Song details in JSON format, including:  
      - `releaseDate` (string): Release date of the song (e.g., "16.07.2006").  
      - `text` (string): Song lyrics.  
      - `link` (string): Link to the song (e.g., YouTube).  
    - **400 Bad Request**: Invalid request.  
    - **500 Internal Server Error**: Internal server error.  

## Installation & Running the Application  

1. **Clone the repository:**  

   ```bash
   git clone https://github.com/mikhail-kirillov/Music-info.git
   ```

2. **Navigate to the project directory:**  

   ```bash
   cd Music-info
   ```

3. **Set up the `.env` file:**  

   Create a `.env` file in the root directory and specify the required configuration parameters. Example content:  

   ```env
    POSTGRES_USER=your_db_user
    POSTGRES_PASSWORD=your_db_password
    POSTGRES_PORT=your_db_port
    POSTGRES_DB=your_db_name
    SERVER_PORT=your_server_port
    MUSIC_API_URL=your_url
   ```

4. **Run the application using `make`:**  

   ```bash
   make
   ```

   This command will build and start the application in Docker containers.  

5. **Access the server:**  

   Once successfully started, the server will be available at the configured address (e.g., `http://localhost:8080`). You can use tools like `curl` or Postman to interact with the API.  

## API Documentation  

After launching the application, the API documentation generated using `swaggo` will be available at `http://localhost:8080/swagger/index.html`.
