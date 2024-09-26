# Groupie tracker

## Project Overview

**Groupie tracker** is a web application built to display information about various artists and bands. It consumes data from a provided API, visualizes this data through a user-friendly interface, and includes client-server interactions for enhanced functionality. The project successfully displays artist information, concert locations, concert dates, and the relationships between these entities.

## API Structure

The API consists of four main parts:

1. **Artists**:
   - Contains information such as:
     - Name
     - Image
     - Year they began their activity
     - Date of their first album
     - Band members

2. **Locations**:
   - Displays the last and/or upcoming concert locations for each artist or band.

3. **Dates**:
   - Displays the last and/or upcoming concert dates for each artist or band.

4. **Relation**:
   - Links the artists, dates, and locations together for easy cross-referencing.

## Features

- **Data Visualization**: 
  - Information about artists is displayed using various visual formats like:
    - Cards
    - Tables
    - Lists
    - Graphs
  - Each format is designed for easy navigation and a pleasant user experience.
  
- **Artist Profiles**:
  - Displays detailed information about each artist or band, including their history, members, and first album.

- **Concert Information**:
  - Provides upcoming and past concert locations and dates for each band or artist.
  - Interactive maps and date lists help users explore concert details.

- **Client-Server Interaction**:
  - Implemented event-based actions, allowing real-time data updates via client-server requests.
  - Features such as filtering artists by location or date trigger server requests and dynamically update the displayed information.
  - Follows the [request-response](https://en.wikipedia.org/wiki/Request%E2%80%93response) pattern for client-server communication.

- **Responsive Design**:
  - The website is fully responsive and works across various screen sizes, ensuring a consistent experience on both desktop and mobile devices.

## Event-Driven Interactions

The project incorporates dynamic events where the user can trigger server requests, such as:

- **Filtering by Concert Location**: Users can filter artists based on upcoming or past concert locations, sending requests to the server and dynamically updating the list of artists.
- **Real-time Concert Updates**: The website automatically fetches real-time concert data when triggered by specific user actions (e.g., selecting a band, location, or date).
  
These interactions enhance the user experience by providing live updates and reducing the need for manual refreshes.

## Technology Stack

- **Frontend**: HTML, CSS, JavaScript (with data visualization libraries)
- **Backend**: Go (for server-side requests and handling API calls)
- **API**: Custom API providing artists, locations, dates, and relation data
- **Client-Server Communication**: RESTful architecture, following the request-response pattern

## How to Run the Project

1. Clone the repository:
    ```bash
    git clone   https://github.com/kevwasonga/groupie-tracker
    cd groupie-tracker
    ```

2. Install dependencies and run the server:
    ```bash
    go run main.go
    ```

3. Open your browser and navigate to:
    ```bash
    http://localhost:8080
    ```

4. Explore the various sections and features of the website.

## Future Enhancements

Some potential future improvements include:

- **Search Functionality**: Add a search bar to quickly find artists by name, location, or date.
- **User Accounts**: Allow users to create accounts and save their favorite bands or concerts.
- **Notification System**: Notify users of upcoming concerts based on their preferences.

## Conclusion

This project demonstrates the integration of API data, client-server interaction, and data visualization to create a fully functioning web app. Users can explore artist information and concert details seamlessly, with real-time data updates.

---

Thank you for using **Groupie tracker**! We hope you enjoy exploring the music world through our platform.

## Contributors

This is part of 01-edu cirriculum project members were
- [Kevin Wasonga](linkedin.com/in/kevin-wasonga-3a9050317)
- [Maina Anne](linkedin.com/in/maina-anne-37797820b)
- [Granton Onyango](linkedin.com/in/granton-onyango-298ba6213)

Feel free to add a contribution to the project.

## LICENSE
The project is licensed under the [MIT LICENSE](LICENSE)
