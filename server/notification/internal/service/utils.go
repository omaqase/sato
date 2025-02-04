package service

import (
	"strings"
	"text/template"
)

func GenerateFinalHTML(content, username, supportLink string) (string, error) {
	const wrapperTemplate = `
		<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Sato Notification</title>
    <style>
        body {
            font-family: 'Arial', sans-serif;
            background-color: #ffffff;
            margin: 0;
            padding: 0;
        }
		a {
			color: white;
		}
        .container {
            max-width: 600px;
            margin: 20px auto;
            background-color: #ffffff;
            border-radius: 8px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
            overflow: hidden;
        }
        .header {
            background-color: #424242;
            color: #ffffff;
            text-align: center;
            padding: 20px;
            font-size: 20px;
            font-weight: 400;
        }
        .content {
            padding: 32px;
            line-height: 1.6;
            color: #424242;
        }
        .content h1 {
            font-size: 20px;
            margin-bottom: 24px;
            font-weight: 400;
        }
        .content p {
            margin-bottom: 24px;
        }
        .footer {
            background-color: #fafafa;
            text-align: center;
            padding: 16px;
            font-size: 12px;
            color: #666666;
        }
        .message {
            margin-bottom: 24px;
        }
        .button {
            display: inline-block;
            background-color: #424242;
            color: #0066FF;
            text-decoration: none;
            padding: 12px 24px;
            border-radius: 4px;
            font-size: 14px;
            transition: all 0.2s ease;
        }
        .button:hover {
            background-color: #383838;
        }
        @media (max-width: 600px) {
            .container {
                margin: 10px;
                border-radius: 4px;
            }
            .content {
                padding: 24px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            Sato Notification
        </div>
        <div class="content">
            <h1>Hello, {{.Username}}!</h1>
            <div class="message">{{.Content}}</div>
            <a href="{{.SupportLink}}" class="button">Contact Support</a>
        </div>
        <div class="footer">
            &copy; 2025 Sato Marketplace. All rights reserved.
        </div>
    </div>
</body>
</html>
    `

	data := map[string]string{
		"Username":    username,
		"Content":     content,
		"SupportLink": supportLink,
	}

	tmpl, err := template.New("final").Parse(wrapperTemplate)
	if err != nil {
		return "", err
	}

	var finalHTML strings.Builder
	err = tmpl.Execute(&finalHTML, data)
	if err != nil {
		return "", err
	}

	return finalHTML.String(), nil
}
