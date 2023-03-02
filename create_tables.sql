DROP DATABASE IF EXISTS chicagobi;
CREATE DATABASE chicagobi;

DROP TABLE IF EXISTS buildingPermit;
DROP TABLE IF EXISTS ccvi;
DROP TABLE IF EXISTS censusData;
DROP TABLE IF EXISTS communityBound;
DROP TABLE IF EXISTS covidCasesZip;
DROP TABLE IF EXISTS covidDaily;
DROP TABLE IF EXISTS pubHealthStats;
DROP TABLE IF EXISTS taxiTrips;
DROP TABLE IF EXISTS transNetProviderTrips;

\c chicagobi;

CREATE TABLE buildingPermit (

    ID                  INT,    
	PermitNumber        VARCHAR(100),
	PermitType          VARCHAR(100),
	PermitSubType       VARCHAR(100),
	PermitDescription   VARCHAR(100),
	PermitIssueDate     VARCHAR(100),
	PermitEstimatedCost INT,
	PermitStatus        VARCHAR(100),
	PermitStreetNumber  INT,
	PermitStreetName    VARCHAR(100),
	PermitSuffix        VARCHAR(100),
	PermitWorkType      VARCHAR(100),
	PermitPIN1          INT,
	PermitPIN2          INT
);

CREATE TABLE ccvi (

	Index                INT,
	CommunityArea        VARCHAR(100),
	OverallScore         FLOAT, 
	SocioeconomicScore   FLOAT, 
	HouseholdCrowding    FLOAT, 
	NoVehicleHouseholds  FLOAT, 
	PerCapitaIncome      FLOAT, 
	Unemployment         FLOAT, 
	NoHighSchoolDiploma  FLOAT, 
	AgeAdjustedDeathRate FLOAT, 
	DiabetesPrevalence   FLOAT, 
	HIVPrevalenceRate    FLOAT, 
	InfantMortalityRate  FLOAT, 
	LeadPoisoningRate    FLOAT, 
	HospitalizationRate  FLOAT
);

CREATE TABLE censusData (

	CommunityAreaNumber                   INT,
	CommunityAreaName                     VARCHAR(100),
	PercentHouseholdsBelowPoverty         FLOAT, 
	PercentAged16Unemployed               FLOAT, 
	PercentAged25WithoutHighSchoolDiploma FLOAT, 
	PercentAgedUnder18OrOver64            FLOAT, 
	PerCapitaIncome                       FLOAT, 
	HardshipIndex                         INT
);

CREATE TABLE communityBound (

	AreaNumber                            INT,
	AreaName                              VARCHAR(100),
	Geometry                              VARCHAR(100) 
	
);

CREATE TABLE covidCasesZip (

    Zip        VARCHAR(100),
    Cases      INT,
    Tests      INT,
    Deaths     INT,
    PeopleTest INT

);

CREATE TABLE covidDaily (

	Date                VARCHAR(20),
	TotalCases          VARCHAR(20),
	NewCases            VARCHAR(20),
	TotalDeaths         VARCHAR(20),
	NewDeaths           VARCHAR(20),
	TotalHospitalized   VARCHAR(20),
	NewHospitalized     VARCHAR(20),
	AverageDailyHospitalized VARCHAR(20)

);

CREATE TABLE pubHealthStats (

    CommunityArea      VARCHAR(20),
    BirthRate          VARCHAR(20),
    GeneralFertility   VARCHAR(20),
    LowBirthWeight     VARCHAR(20),
    PrenatalCare       VARCHAR(20),
    TeenBirthRate      VARCHAR(20),
    Uninsured          VARCHAR(20),
    BelowPoverty       VARCHAR(20),
    CrowdedHousing     VARCHAR(20),
    Dependency         VARCHAR(20),
    NoDiploma          VARCHAR(20),
    PerCapitaIncome    VARCHAR(20),
    Unemployment       VARCHAR(20),
    Assault            VARCHAR(20),
    BreastCancer       VARCHAR(20),
    Cancer             VARCHAR(20),
    ColorectalCancer   VARCHAR(20),
    Diabetes           VARCHAR(20),
    FirearmMortality   VARCHAR(20),
    InfantMortality    VARCHAR(20),
    LungCancer         VARCHAR(20),
    ProstateCancer     VARCHAR(20),
    Stroke             VARCHAR(20),
    Tuberculosis       VARCHAR(20),
    BelowPovertyRecent VARCHAR(20),
    NoDiplomaRecent    VARCHAR(20),
    UnemploymentRecent VARCHAR(20)
);

CREATE TABLE taxiTrips (

	ID                   INT,
	TripStartTimestamp   TIMESTAMP,
	TripEndTimestamp     TIMESTAMP,
	TripSeconds          INT,
	TripMiles            FLOAT, 
	PickupCensusTract    VARCHAR(20),
	DropoffCensusTract   VARCHAR(20),
	PickupCommunityArea  VARCHAR(20),
	DropoffCommunityArea VARCHAR(20),
	Fare                 FLOAT, 
	Tips                 FLOAT, 
	Tolls                FLOAT, 
	Extras               FLOAT, 
	TripTotal            FLOAT, 
	PaymentType          VARCHAR(20),
	Company              VARCHAR(20),
	TaxiID               VARCHAR(20)

);

CREATE TABLE transNetProviderTrips (

	ID                         INT,
	TripID                     VARCHAR(20),
	PickupCommunityArea        VARCHAR(20),
	DropoffCommunityArea       VARCHAR(20),
	TripStartTimestamp         time.Time
	TripEndTimestamp           time.Time
	TripSeconds                INT,
	TripMiles                  FLOAT, 
	PickupCensusTract          VARCHAR(20),
	DropoffCensusTract         VARCHAR(20),
	PickupCentroidLatitude     FLOAT, 
	PickupCentroidLongitude    FLOAT, 
	DropoffCentroidLatitude    FLOAT, 
	DropoffCentroidLongitude   FLOAT, 
	SharedTripAuthorized       VARCHAR(20),
	TripsPooled                INT,
	PickupCommunityName        VARCHAR(20),
	DropoffCommunityName       VARCHAR(20),
	Fare                       FLOAT, 
	Tip                        FLOAT, 
	AdditionalCharges          FLOAT, 
	TripTotal                  FLOAT, 
	SharedTripPaymentType      VARCHAR(20),
	PaymentType                VARCHAR(20),
	Company                    VARCHAR(20),
	TripType                   VARCHAR(20),
	TripExtras                 FLOAT, 
	TripStartTimeAdjusted      time.Time
	TripEndTimeAdjusted        time.Time
	TripMilesAdjusted          FLOAT, 
	TripFareAdjusted           FLOAT, 
	TripTotalAdjusted          FLOAT, 
	SharedTripActualDistance   FLOAT, 
	SharedTripMatchedID        VARCHAR(20),
	OriginationLatitude        FLOAT, 
	OriginationLongitude       FLOAT, 
	DestinationLatitude        FLOAT, 
	DestinationLongitude       FLOAT, 
	SharedTripCost             FLOAT, 
	NumberOfMatchedSharedTrips INT

);