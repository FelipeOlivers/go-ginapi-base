[![GNU License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<br />
<p align="center">
    <a href="#">
      <img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" alt="Logo" width="80" height="120">
    </a>

  <h3 align="center">GO GinAPI Base</h3>

  <p align="center">
    Base to start new project in Golang with Gin
  </p>
</p>

<!-- ABOUT THE PROJECT -->
## About The Project

:smile: This project aims to provide an initial base to start developing a new API with GO using the Gin framework. 



### Built With

* [GO](https://golang.org/)
* [Gin Framework](https://github.com/gin-gonic/gin)
* [GORM](https://gorm.io/)
* [PostgreSQL](https://www.postgresql.org/)



<!-- GETTING STARTED -->
## Getting Started

This is an example for you to start an API written in GO using the Gin framework.
To get a working local copy, follow these simple example steps.

### Prerequisites

* Docker
  ```sh
  docker --version
  ```

### Installation / Usage

1. Clone the repo
   ```sh
   git clone https://github.com/felipeolivers/go-ginapi-base.git
   ```
2. Execute docker command to build
   ```sh
   sudo docker build . -t go-ginapi
   ```

3. Execute docker command to run application
   ```sh
   sudo docker run -i -t -d -p 8000:8000 go-ginapi
   ```

4. Enter your API in `Browser`
   ```JS
    http://localhost:8000/api/v1/users/1
   ```



<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/felipeolivers/go-ginapi-base) for a list of proposed features (and known issues).



<!-- LICENSE -->
## License

Distributed under the GNU General Public License v3.0. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Luiz Filipe Miranda de Oliveira - [@felipeolivers](https://www.instagram.com/felipeolivers/) - folivers@gmail.com

Project Link: [https://github.com/FelipeOlivers/go-ginapi-base](https://github.com/FelipeOlivers/go-ginapi-base)



<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [GO](https://golang.org/)
* [Gin Framework](https://github.com/gin-gonic/gin)
* [GORM](https://gorm.io/)
* [PostgreSQL](https://www.postgresql.org/)
* [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
* [GitHub Pages](https://pages.github.com)
* [Font Awesome](https://fontawesome.com)



<!-- MARKDOWN LINKS & IMAGES -->
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/felipeolivers/go-ginapi-base/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/folivers
[product-screenshot]: images/screenshot.png