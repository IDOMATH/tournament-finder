import { Component } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { Tournament } from "../tournament.model";

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
  enteredIsBoysVarsity = false;
  enteredIsGirlsVarsity = false;
  enteredIsBoysJv = false;
  enteredIsGirlsJv = false;
  enteredIsBoysMs = false;
  enteredIsGirlsMs = false;
  enteredIsBoysYouth = false;
  enteredIsGirlsYouth = false;

  onSubmit() {
    let tournament: Tournament = {
      name: this.enteredTournamentName,
      locationName: this.enteredLocationName,
    };
    console.log("submitting");
    console.log(this.enteredIsBoysVarsity);
  }
}
