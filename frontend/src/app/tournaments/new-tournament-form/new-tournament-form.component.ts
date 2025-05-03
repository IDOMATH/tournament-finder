import { Component } from "@angular/core";
import {
  FormArray,
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
    streetAddress: new FormControl("", {}),
    city: new FormControl("", {}),
    state: new FormControl("", {}),
    dates: new FormGroup({
      startDate: new FormControl("", {}),
      endDate: new FormControl("", {}),
    }),
    ageDivision: new FormGroup({
      isBoysVarsity: new FormControl(false, {}),
      isGirlsVarsity: new FormControl(false, {}),
      isBoysJv: new FormControl(false, {}),
      isGirlsJv: new FormControl(false, {}),
      isBoysMs: new FormControl(false, {}),
      isGirlsMs: new FormControl(false, {}),
      isBoysYouth: new FormControl(false, {}),
      isGirlsYouth: new FormControl(false, {}),
    }),
  });

  onSubmit() {
    console.log("submitting");
    console.log(this.form.controls.name);
  }
}
