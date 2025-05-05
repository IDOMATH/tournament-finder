import { Component, inject } from "@angular/core";
import {
  FormArray,
  FormControl,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from "@angular/forms";
import { Tournament } from "../tournament.model";
import { TournamentService } from "../tournament-service";

@Component({
  selector: "app-new-tournament-form",
  imports: [ReactiveFormsModule],
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
    dates: new FormGroup({
      startDate: new FormControl("", { validators: [Validators.required] }),
      endDate: new FormControl("", { validators: [Validators.required] }),
    }),
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
    const tournament = {
      name: this.form.controls.name.value!,
      locationName: this.form.controls.locationName.value!,
      streetAddress: this.form.controls.streetAddress.value!,
      city: this.form.controls.city.value!,
      state: this.form.controls.state.value!,
      startDate: new Date(this.form.controls.dates.controls.startDate.value!),
      endDate: new Date(this.form.controls.dates.controls.endDate.value!),
      isBoysVarsity:
        this.form.controls.ageDivision.controls.isBoysVarsity.value!,
      isGirlsVarsity:
        this.form.controls.ageDivision.controls.isGirlsVarsity.value!,
      isBoysJv: this.form.controls.ageDivision.controls.isBoysJv.value!,
      isGirlsJv: this.form.controls.ageDivision.controls.isGirlsJv.value!,
      isBoysMs: this.form.controls.ageDivision.controls.isBoysMs.value!,
      isGirlsMs: this.form.controls.ageDivision.controls.isGirlsMs.value!,
      isBoysYouth: this.form.controls.ageDivision.controls.isBoysYouth.value!,
      isGirlsYouth: this.form.controls.ageDivision.controls.isGirlsYouth.value!,
      id: 0,
      organizerId: 0,
      isFull: false,
    };

    console.log("submitting");
    console.log(this.form.controls.name);
    this.form.controls.ageDivision.controls.isBoysJv;
    this.tournamentService.postTournament(tournament);
  }
}
