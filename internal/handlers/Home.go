package handlers

import "github.com/gin-gonic/gin"

func Home(c *gin.Context) {
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>My Web Server</title>
		<style>
			body { 
				font-family: Arial, sans-serif; 
				margin: 0; 
				padding: 40px; 
				min-height: 100vh;
			}
			.container { 
				background: white; 
				padding: 40px; 
				border-radius: 15px; 
				box-shadow: 0 10px 30px rgba(0,0,0,0.2);
				max-width: 800px;
				margin: 0 auto;
			}
			h1 { 
				color: #333; 
				text-align: center; 
				font-size: 2.5rem;
				margin-bottom: 10px;
			}
			p { 
				color: #666; 
				text-align: center; 
				font-size: 1.2rem;
				line-height: 1.6;
			}
			.links { 
				text-align: center; 
				margin-top: 30px; 
			}
			a { 
				background: #667eea;
				color: white;
				padding: 12px 25px;
				text-decoration: none; 
				border-radius: 25px;
				margin: 0 10px;
				display: inline-block;
				transition: all 0.3s ease;
			}
			a:hover { 
				background: #764ba2;
				transform: translateY(-2px);
				box-shadow: 0 5px 15px rgba(0,0,0,0.2);
			}
			.emoji { font-size: 2rem; }
		</style>
	</head>
	<body>
		<div class="container">
		
			<h1>Welcome to My Web Server!</h1>
			<p>This is a modern HTTP server built with <strong>Go</strong> and <strong>Gin framework</strong>.</p>
			<p>Server is running successfully on port <strong>:8080</strong></p>
		</div>
	</body>
	</html>`

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, html)
}
