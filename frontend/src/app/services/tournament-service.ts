import { Injectable } from "@angular/core";
import { Tournament } from "../models/tournament.model";

@Injectable({ providedIn: "root" })
export class TournamentService {
  getTournamentById(id: number): Tournament {
    const tournament = {
      name: "Test Tournament",
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
      id: id,
      organizerId: 0,
    };

    return tournament;
  }
  getTournaments(): Tournament[] {
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
  }
  putTournament(tournament: Tournament) {
    JSON.stringify(tournament);
  }
  deleteTournament(id: number) {}
}
