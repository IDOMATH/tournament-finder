import { Component, inject } from "@angular/core";
import {
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from "@angular/forms";
import { TournamentService } from "../../services/tournament-service";
import { StatePickerComponent } from "../../state-picker/state-picker.component";
import { endDateNotBeforeStartDate } from "../../validators/end-date-after-begin-date";

@Component({
  selector: "app-new-tournament-form",
  imports: [ReactiveFormsModule, StatePickerComponent],
  templateUrl: "./new-tournament-form.component.html",
  styleUrl: "./new-tournament-form.component.css",
})
export class NewTournamentFormComponent {
  private tournamentService = inject(TournamentService);

  form = new FormGroup({
    name: new FormControl("", { validators: [Validators.required] }),
    locationName: new FormControl("", { validators: [Validators.required] }),
    streetAddress: new FormControl("", { validators: [Validators.required] }),
    city: new FormControl("", { validators: [Validators.required] }),
    state: new FormControl("", { validators: [Validators.required] }),
    dates: new FormGroup(
      {
        startDate: new FormControl("", { validators: [Validators.required] }),
        endDate: new FormControl("", { validators: [Validators.required] }),
      },
      { validators: [endDateNotBeforeStartDate("startDate", "endDate")] }
    ),
    ageDivision: new FormGroup({
      isBoysVarsity: new FormControl(false, {
        validators: [Validators.required],
      }),
      isGirlsVarsity: new FormControl(false, {
        validators: [Validators.required],
      }),
      isBoysJv: new FormControl(false, { validators: [Validators.required] }),
      isGirlsJv: new FormControl(false, { validators: [Validators.required] }),
      isBoysMs: new FormControl(false, { validators: [Validators.required] }),
      isGirlsMs: new FormControl(false, { validators: [Validators.required] }),
      isBoysYouth: new FormControl(false, {
        validators: [Validators.required],
      }),
      isGirlsYouth: new FormControl(false, {
        validators: [Validators.required],
      }),
    }),
  });

  onSubmit() {
    if (this.form.invalid) {
      return;
    }
    const tournament = {
      name: this.form.controls.name.value!,
      locationName: this.form.controls.locationName.value!,
      streetAddress: this.form.controls.streetAddress.value!,
      city: this.form.controls.city.value!,
      state: this.form.controls.state.value!,
      startDate: new Date(this.form.controls.dates.controls.startDate.value!),
      endDate: new Date(this.form.controls.dates.controls.endDate.value!),
      boysVarsity: this.form.controls.ageDivision.controls.isBoysVarsity.value!
        ? 1
        : 0,
      girlsVarsity: this.form.controls.ageDivision.controls.isGirlsVarsity
        .value!
        ? 1
        : 0,
      boysJv: this.form.controls.ageDivision.controls.isBoysJv.value! ? 1 : 0,
      girlsJv: this.form.controls.ageDivision.controls.isGirlsJv.value! ? 1 : 0,
      boysMs: this.form.controls.ageDivision.controls.isBoysMs.value! ? 1 : 0,
      girlsMs: this.form.controls.ageDivision.controls.isGirlsMs.value! ? 1 : 0,
      boysYouth: this.form.controls.ageDivision.controls.isBoysYouth.value!
        ? 1
        : 0,
      girlsYouth: this.form.controls.ageDivision.controls.isGirlsYouth.value!
        ? 1
        : 0,
      id: 0,
      organizerId: 0,
    };

    console.log("submitting");
    console.log(this.form.controls.name);
    this.form.controls.ageDivision.controls.isBoysJv;
    this.tournamentService.postTournament(tournament);
  }
}
