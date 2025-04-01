import { Component } from "@angular/core";
import { FormsModule } from "@angular/forms";

@Component({
  selector: "app-new-tournament-form",
  imports: [FormsModule],
  templateUrl: "./new-tournament-form.component.html",
  styleUrl: "./new-tournament-form.component.css",
})
export class NewTournamentFormComponent {
  enteredTournamentName = "";
  enteredLocationName = "";
  enteredStreetAddress = "";
  enteredCity = "";
  enteredState = "";
  enteredStartDate = "";
  enteredEndDate = "";

  onSubmit() {
    console.log("submitting");
  }
}
