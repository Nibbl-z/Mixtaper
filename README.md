![Mixtaper](https://github.com/Nibbl-z/Mixtaper/blob/main/client/static/logo.png?raw=true)

### A website to upload and share custom levels for [Bits and Bops!](https://store.steampowered.com/app/1929290/Bits__Bops/)

You can post and download levels, search for songs, and customize your profile!

![image](https://github.com/user-attachments/assets/8f417eb9-0fa3-4857-b86e-61ab38d0b937)


# Features

- Posting levels
- Downloading levels
- Level search
- Automatically get data from uploaded level, such as games used and BPM
- Account customization (About me, display name, and profile picture)
- Account settings

# Technologies Used

- Frontend: TypeScript and Svelte
- Backend: Go with Atreugo
- Storage, Database, and Authentication: Appwrite

# Self-Hosting (Frontend)

1. Install NodeJS
2. Put your backend URL (if you do not have one and wish to only make frontend changes, use `https://mixtaperbackend.nibbles.hackclub.app` and Appwrite project ID (if you once again do not have one, use `676f205d000370a15786`) into `.env.dist`, and rename it to `.env`
3. Run `npm run dev`

# Self-Hosting (Backend)

1. Create an [Appwrite](https://appwrite.io/) account, and create a project
2. In the overview tab, go to the bottom of the page, select API Keys, then click Create API Key
3. Give it a name and give it all scopes
4. Go to your API key and copy the secret
5. Go into `/playgrounds/.env.dist`, and fill in your Appwrite endpoint (which is `https://cloud.appwrite.io/v1` by default), Appwrite project ID, and your API key, then rename it to `.env`
6. Install Go
7. Run `go run .` in `/playgrounds`. This will set up your Appwrite project to have everything necessary for the backend to function
8. After the setup is complete, copy the `.env` from `/playgrounds` and put it into `/server`
9. In `/server/main.go`, change the `Addr` on line 17 to the URL/port you wish to use
10. In `/server/main.go`, change the `Access-Control-Allow-Origin` on lines 21 and 29 to the URL that your frontend will be hosted on
11. Run `go run .` in `/server`
