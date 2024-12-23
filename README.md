# Create-a-simple-web-server---Using-Go
Basic way to create the server with Go

## Features

1. **Dynamic Routing**:
   - Example route `/hello` that responds with `"Hello, World!"`.
   
2. **Static File Serving**:
   - Serves static files from the `public/` directory.
   - Example: Access the `index.html` file at `/`.

---

## Setup and Running the Code

Follow these steps to run the project locally:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/Avinashvstudio/Create-a-simple-web-server---Using-Go.git
   cd Create-a-simple-web-server---Using-Go


Here’s how you can push your project to GitHub and create a `README.md` file with explanations and instructions:

---

### **Step 1: Push Your Code to GitHub**

1. **Navigate to Your Project Directory**:
   ```bash
   cd path/to/web-server
   ```

2. **Initialize Git**:
   ```bash
   git init
   ```

3. **Add Remote Repository**:
   Link your local repository to the GitHub repository:
   ```bash
   git remote add origin https://github.com/Avinashvstudio/Create-a-simple-web-server---Using-Go.git
   ```

4. **Stage Your Files**:
   Add all files to the staging area:
   ```bash
   git add .
   ```

5. **Commit Your Changes**:
   Create a commit with a descriptive message:
   ```bash
   git commit -m "Initial commit: Simple Go web server"
   ```

6. **Push to GitHub**:
   Push the code to the `main` branch of your GitHub repository:
   ```bash
   git branch -M main
   git push -u origin main
   ```

---

### **Step 2: Create a `README.md` File**

Create a `README.md` file in the root of your project directory. Add the following content:

```markdown
# Simple Web Server in Go

This project demonstrates how to create a simple web server using the Go programming language. It handles dynamic routes and serves static files.

---

## File Structure

```
web-server/
├── main.go            # Main entry point for the server
├── handlers/          # Directory for handler logic
│   └── hello.go       # Example handler
├── public/            # Directory for static assets (HTML, CSS, JS, images, etc.)
│   └── index.html     # Example HTML file
└── go.mod             # Module file for Go
```

---

## Features

1. **Dynamic Routing**:
   - Example route `/hello` that responds with `"Hello, World!"`.
   
2. **Static File Serving**:
   - Serves static files from the `public/` directory.
   - Example: Access the `index.html` file at `/`.

---

## Setup and Running the Code

Follow these steps to run the project locally:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/Avinashvstudio/Create-a-simple-web-server---Using-Go.git
   cd Create-a-simple-web-server---Using-Go
   ```

2. **Install Go**:
   Ensure Go is installed on your machine. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

3. **Run the Server**:
   ```bash
   go run main.go
   ```

4. **Access the Web Server**:
   - Open your browser and navigate to:
     - `http://localhost:8080/` for the static homepage.
     - `http://localhost:8080/hello` for the dynamic route.

---

## Example Outputs

### **Static Route (`/`)**
When accessing `/`, the `index.html` file from the `public/` directory is served:
```
Welcome to the Go Web Server!
Try accessing /hello.
```

### **Dynamic Route (`/hello`)**
When accessing `/hello`, the server responds with:
```
Hello, World!
```

---

## Contributing

Feel free to fork this repository and submit pull requests for enhancements or bug fixes.

---

## License

This project is open-source and available under the MIT License.
```

---

### **Step 3: Push the `README.md` File**

1. **Stage the `README.md` File**:
   ```bash
   git add README.md
   ```

2. **Commit the Changes**:
   ```bash
   git commit -m "Added README.md with project details"
   ```

3. **Push the Changes**:
   ```bash
   git push
   ```

---

### **Step 4: Verify the Repository**

- Go to your GitHub repository at `https://github.com/Avinashvstudio/Create-a-simple-web-server---Using-Go`.
- Confirm that the files and `README.md` are uploaded correctly.

---
