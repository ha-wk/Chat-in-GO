Real-Time Chat Application with Go, Fiber, and HTMX
This project demonstrates how to build a real-time chat application using Go, the Fiber web framework, and HTMX for front-end interactivity. The application supports real-time updates using WebSockets, allowing users to chat instantly without page reloads.

Features
Real-time messaging using WebSockets
Lightweight and fast with Go and Fiber
Minimal front-end code with HTMX for seamless interactivity
Simple and clean UI with Bootstrap
Getting Started
Follow the instructions below to set up and run the project on your local machine.

Prerequisites
Ensure you have the following installed on your system:

Go (Golang) (version 1.18 or later)
A code editor, such as Visual Studio Code
Git for version control
Installation
Clone the repository:

bash
Copy code
git clone https://github.com/yourusername/real-time-chat-app.git
cd real-time-chat-app
Install dependencies:

Ensure you have Fiber and html/template for template rendering:

go get -u github.com/gofiber/fiber/v2
go get -u github.com/gofiber/template/html/v2
Set up the project structure:

The project structure should look like this:

real-time-chat-app/
├── main.go
├── views/
│   ├── index.html
│   └── chat.html
├── static/
│   └── style.css
└── README.md
Project Structure
main.go: The main entry point of the application where routes, WebSocket connections, and server configurations are defined.
views/: Directory containing HTML templates.
index.html: The homepage or entry point to the chat application.
chat.html: The main chat interface with WebSocket connections.
static/: Directory for static files such as CSS.
style.css: Custom styles for the chat application.
Running the Application
Start the Go server:

In the root directory of the project, run the following command:

go run main.go
Open your browser:

Visit http://localhost:3000 in your web browser to access the chat application.

How It Works
WebSocket Connection: The server establishes a WebSocket connection with each client using the HTMX extension for WebSockets. This allows real-time bidirectional communication.

Real-Time Updates: When a user sends a message, it's broadcast to all connected clients in real-time using the WebSocket connection.

Dynamic UI with HTMX: The front-end leverages HTMX to manage real-time updates without requiring a JavaScript-heavy framework. HTMX handles WebSocket messages and dynamically updates the DOM.

Built With
Go - The Go programming language for backend development.
Fiber - An Express-inspired web framework for Go.
HTMX - A JavaScript library that simplifies dynamic web applications.
Bootstrap - A CSS framework for responsive design.
Contributing
Contributions are welcome! If you have suggestions or find issues, feel free to open an issue or submit a pull request.