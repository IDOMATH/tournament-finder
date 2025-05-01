import { Component } from "@angular/core";
import {
  FormControl,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
} from "@angular/forms";
import { Tournament } from "../tournament.model";

@Component({
  selector: "app-new-tournament-form",
  imports: [ReactiveFormsModule],
  templateUrl: "./new-tournament-form.component.html",
  styleUrl: "./new-tournament-form.component.css",
})
export class NewTournamentFormComponent {
  form = new FormGroup({
    name: new FormControl("", {}),
    locationName: new FormControl("", {}),
  });
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
    console.log("submitting");
    console.log(this.enteredIsBoysVarsity);
  }
}
