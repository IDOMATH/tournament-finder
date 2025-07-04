import { Component, output } from "@angular/core";

// type StateOptions =
//   | "Alabama"
//   | "Alaska"
//   | "Arizona"
//   | "Arkansas"
//   | "California"
//   | "Colorado"
//   | "Connecticut"
//   | "Delaware"
//   | "Florida"
//   | "Georgia"
//   | "Hawaii"
//   | "Idaho"
//   | "Illinois"
//   | "Indiana"
//   | "Iowa"
//   | "Kansas"
//   | "Kentucky"
//   | "Louisiana"
//   | "Maine"
//   | "Maryland"
//   | "Massachusetts"
//   | "Michigan"
//   | "Minnesota"
//   | "Mississippi"
//   | "Missouri"
//   | "Minnesota"
//   | "Mississippi"
//   | "Missouri"
//   | "MontanaNebraska"
//   | "Nevada"
//   | "New Hampshire"
//   | "New Jersey"
//   | "New Mexico"
//   | "New York"
//   | "North Carolina"
//   | "North Dakota"
//   | "Ohio"
//   | "Oklahoma"
//   | "Oregon"
//   | "Pennsylvania"
//   | "Rhode Island"
//   | "South Carolina"
//   | "South Dakota"
//   | "Tennessee"
//   | "Texas"
//   | "Utah"
//   | "Vermont"
//   | "Virginia"
//   | "Washington"
//   | "West Virginia"
//   | "Wisconsin"
//   | "Wyoming";

const stateOptions = [
  "Alabama",
  "Alaska",
  "Arizona",
  "Arkansas",
  "California",
  "Colorado",
  "Connecticut",
  "Delaware",
  "Florida",
  "Georgia",
  "Hawaii",
  "Idaho",
  "Illinois",
  "Indiana",
  "Iowa",
  "Kansas",
  "Kentucky",
  "Louisiana",
  "Maine",
  "Maryland",
  "Massachusetts",
  "Michigan",
  "Minnesota",
  "Mississippi",
  "Missouri",
  "Minnesota",
  "Mississippi",
  "Missouri",
  "MontanaNebraska",
  "Nevada",
  "New Hampshire",
  "New Jersey",
  "New Mexico",
  "New York",
  "North Carolina",
  "North Dakota",
  "Ohio",
  "Oklahoma",
  "Oregon",
  "Pennsylvania",
  "Rhode Island",
  "South Carolina",
  "South Dakota",
  "Tennessee",
  "Texas",
  "Utah",
  "Vermont",
  "Virginia",
  "Washington",
  "West Virginia",
  "Wisconsin",
  "Wyoming",
];

@Component({
  selector: "app-state-picker",
  imports: [],
  templateUrl: "./state-picker.component.html",
  styleUrl: "./state-picker.component.css",
})
export class StatePickerComponent {
  selectedState = output<string>();
  stateOptions = [
    "Alabama",
    "Alaska",
    "Arizona",
    "Arkansas",
    "California",
    "Colorado",
    "Connecticut",
    "Delaware",
    "Florida",
    "Georgia",
    "Hawaii",
    "Idaho",
    "Illinois",
    "Indiana",
    "Iowa",
    "Kansas",
    "Kentucky",
    "Louisiana",
    "Maine",
    "Maryland",
    "Massachusetts",
    "Michigan",
    "Minnesota",
    "Mississippi",
    "Missouri",
    "Minnesota",
    "Mississippi",
    "Missouri",
    "MontanaNebraska",
    "Nevada",
    "New Hampshire",
    "New Jersey",
    "New Mexico",
    "New York",
    "North Carolina",
    "North Dakota",
    "Ohio",
    "Oklahoma",
    "Oregon",
    "Pennsylvania",
    "Rhode Island",
    "South Carolina",
    "South Dakota",
    "Tennessee",
    "Texas",
    "Utah",
    "Vermont",
    "Virginia",
    "Washington",
    "West Virginia",
    "Wisconsin",
    "Wyoming",
  ];
}
