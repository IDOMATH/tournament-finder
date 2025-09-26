import { inject, Injectable, signal } from "@angular/core";
import { Tournament } from "../models/tournament.model";
import { HttpClient } from "@angular/common/http";
import { map } from "rxjs";

@Injectable({ providedIn: "root" })
export class TournamentService {
  private httpClient = inject(HttpClient);
  tournament = signal<Tournament | undefined>(undefined);

  getTournamentById(id: number) {
    return this.httpClient
      .get<{ tournament: Tournament }>(
        "http://localhost:8080/tournaments/" + id
      )
      .pipe(map((resData) => this.tournament.set(resData.tournament)));
  }
  getTournaments(): Tournament[] {
    this.httpClient.get<{ tournaments: Tournament[] }>(
      "http://localhost:8080/tournaments/"
    );
    const tournaments = [
      {
        name: "Test One",
        locationName: "Joe's House",
        streetAddress: "123 Small St",
        city: "St. Paul",
        state: "Minnesota",
        startDate: new Date(2025, 11, 17),
        endDate: new Date(2025, 11, 17),
        boysVarsity: 2,
        girlsVarsity: 2,
        boysJv: 1,
        girlsJv: 1,
        boysMs: 0,
        girlsMs: 0,
        boysYouth: 0,
        girlsYouth: 0,
        id: 17,
        organizerId: 0,
      },
      {
        name: "Test TOO",
        locationName: "Shakopee High School",
        streetAddress: "123 Small St",
        city: "Shakopee",
        state: "Minnesota",
        startDate: new Date(2025, 11, 17),
        endDate: new Date(2025, 11, 17),
        boysVarsity: 2,
        girlsVarsity: 2,
        boysJv: 1,
        girlsJv: 1,
        boysMs: 0,
        girlsMs: 0,
        boysYouth: 0,
        girlsYouth: 0,
        id: 24,
        organizerId: 0,
      },
    ];

    return tournaments;
  }
  postTournament(tournament: Tournament) {
    JSON.stringify(tournament);
    // call to post tournament
    this.httpClient.post("http://localhost:8080/tournaments/", tournament);
  }
  putTournament(tournament: Tournament) {
    JSON.stringify(tournament);
    this.httpClient.put("http://localhost:8080/tournaments/", tournament);
  }
  deleteTournament(id: number) {
    this.httpClient.delete("http://localhost:8080/tournaments/" + id);
  }
}
