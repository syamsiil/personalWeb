let projectCardData = [];

function projectCard(event) {
  event.preventDefault();

  let projectName = document.getElementById("input-name").value;
  let startDate = new Date(document.getElementById("input-startdate").value);
  let endDate = new Date(document.getElementById("input-enddate").value);
  let description = document.getElementById("input-description").value;
  let image = document.getElementById("input-image").files;

  // alert if endDate>today
  let today = new Date().toISOString().split("T")[0];
  if (endDate > today) {
    return alert("No time travel here!");
  }

  const reactJsIcon =
    '<i class="bx bxl-react mb-3" style="font-size: 32px"></i>';
  const javaScriptIcon =
    '<i class="bx bxl-javascript mb-3" style="font-size: 32px"></i>';
  const nodeJsIcon =
    '<i class="bx bxl-nodejs mb-3" style="font-size: 32px"></i>';
  const css3Icon = '<i class="bx bxl-css3 mb-3" style="font-size: 32px"></i>';

  // alert if no one choose the technologies
  let multiInput = document.querySelectorAll(".multi-input:checked");
  if (multiInput.length === 0) {
    return alert("Select at least one technology use!");
  }

  let reactJsIconDecide = document.getElementById("react-js").checked
    ? reactJsIcon
    : "";
  console.log(reactJsIconDecide);
  let javaScriptIconDecide = document.getElementById("javascript").checked
    ? javaScriptIcon
    : "";
  let nodeJsIconDecide = document.getElementById("node-js").checked
    ? nodeJsIcon
    : "";
  let css3IconDecide = document.getElementById("css3").checked ? css3Icon : "";

  image = URL.createObjectURL(image[0]);
  console.log(image);

  // alert if the  value startdate>enddate
  const sDvalidation = new Date(startDate);
  const eDvalidation = new Date(endDate);
  if (sDvalidation > eDvalidation) {
    return alert("Input your dates correctly!");
  }

  let previewCard = {
    projectName,
    startDate,
    endDate,
    description,
    image,
    reactJsIconDecide,
    javaScriptIconDecide,
    nodeJsIconDecide,
    css3IconDecide,
    postAt: new Date(),
    author: "Seeus",
  };

  projectCardData.push(previewCard);
  console.log(projectCardData);

  renderCard();

  // reset/clear the input field after submit
  document.getElementById("input-name").value = "";
  document.getElementById("input-startdate").value = "";
  document.getElementById("input-enddate").value = "";
  document.getElementById("input-description").value = "";
  document.getElementById("react-js").checked = false;
  document.getElementById("javascript").checked = false;
  document.getElementById("node-js").checked = false;
  document.getElementById("css3").checked = false;
  document.getElementById("input-blog-image").value = "";

  // document.getElementById("project-form").reset();
}

function renderCard() {
  document.getElementById("contents").innerHTML = "";

  for (let index = 0; index < projectCardData.length; index++) {
    const startDate = new Date(projectCardData[index].startDate);
    // console.log(startDate);
    const endDate = new Date(projectCardData[index].endDate);
    // console.log(endDate);
    const differenceTime = endDate - startDate;
    const timeUnits = [
      { value: 365.25 * 24 * 60 * 60 * 1000, label: "years" },
      { value: 30 * 24 * 60 * 60 * 1000, label: "months" },
      { value: 7 * 24 * 60 * 60 * 1000, label: "weeks" },
      { value: 24 * 60 * 60 * 1000, label: "days" },
    ];

    let distanceTime = "";
    for (let calculation = 0; calculation < timeUnits.length; calculation++) {
      const { value, label } = timeUnits[calculation];
      const calculate = Math.floor(differenceTime / value);
      if (calculate > 0) {
        distanceTime = `${calculate} ${label}`;
        break;
      }
    }

    if (distanceTime === "") {
      distanceTime = "today";
    }

    document.getElementById("contents").innerHTML += `
    <div class="list-project-detail" >
        <div class="card" style="width: 18rem">
            <img
                src="${projectCardData[index].image}"
                class="card-img-top"
                alt="project-image"
                style="object-fit: cover; height: 190px"
            />
        <div class="card-body">
             <a
                href="detail-project.html"
                style="
                    font-size: 22px;
                    font-weight: 550;
                    text-decoration: none;
                    color: #000;
                "
                >${projectCardData[index].projectName}</a
             >

        <p style="margin: 0">${getFullTime(
          projectCardData[index].postAt
        )}| Seeus</p>

        <p
                style="
                margin: 0;
                margin-bottom: 10px;
                font-size: 12px;
                color: gray;
                "
        >
        ${distanceTime}
        </p>

        <p class="card-text" style="font-size: 14px">
        ${projectCardData[index].description}
        </p>

        <div class="card-skill">
        ${projectCardData[index].reactJsIconDecide}
        ${projectCardData[index].javaScriptIconDecide}
        ${projectCardData[index].nodeJsIconDecide}
        ${projectCardData[index].css3IconDecide}
        </div>

        <div class="card-button d-flex justify-content-between">
          <a
            href="#"
            class="btn card-btn"
            style="flex: 50%; margin-right: 10px; font-size: 16px"
            >Edit</a
          >
          <a href="#" class="btn card-btn" style="flex: 50%">Delete</a>
        </div>
      </div>
      </div>
    </div>
      `;
  }
}

function getFullTime(time) {
  // console.log("get full time");
  // let time = new Date();
  // console.log(time);

  let monthName = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
  ];
  // console.log(monthName[8]);

  let date = time.getDate();
  // console.log(date);

  let monthIndex = time.getMonth();
  // console.log(monthIndex);

  let year = time.getFullYear();
  // console.log(year);

  let hours = time.getHours();
  let minutes = time.getMinutes();
  // console.log(minutes);

  if (hours <= 9) {
    hours = "0" + hours;
  } else if (minutes <= 9) {
    minutes = "0" + minutes;
  }

  return `${date} ${monthName[monthIndex]} ${year} ${hours}:${minutes} WIB`;
}
